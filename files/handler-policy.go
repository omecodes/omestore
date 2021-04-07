package files

import (
	"context"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/omecodes/store/objects"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"io"
	"net/url"
	"path"
	"strings"

	"github.com/omecodes/errors"
	"github.com/omecodes/libome/logs"
	"github.com/omecodes/store/auth"
	"github.com/omecodes/store/common/cenv"
)

type PolicyHandler struct {
	BaseHandler
}

func (h *PolicyHandler) isAdmin(ctx context.Context) bool {
	user := auth.Get(ctx)
	if user == nil {
		return false
	}
	return user.Name == "admin"
}

func (h *PolicyHandler) assertPermissionIsGranted(ctx context.Context, rules ...string) error {
	var formattedRules []string
	for _, exp := range rules {
		if exp == "true" {
			return nil
		}
		formattedRules = append(formattedRules, "("+exp+")")
	}
	fullExpression := strings.Join(formattedRules, " || ")

	prg, err := cenv.GetProgram(fullExpression,
		cel.Declarations(
			decls.NewVar("user", decls.NewObjectType("User")),
			decls.NewVar("app", decls.NewObjectType("ClientApp")),
			decls.NewVar("data", decls.NewObjectType("Header")),
			decls.NewFunction("now",
				decls.NewOverload(
					"now_uint",
					[]*expr.Type{}, decls.Uint,
				),
			),
		),
		cel.Types(&auth.User{}, &auth.ClientApp{}, &objects.Header{}),
	)
	if err != nil {
		return errors.Internal("context missing access rule evaluator")
	}

	vars := map[string]interface{}{}

	user := auth.Get(ctx)
	if user != nil {
		vars["user"] = user
	}

	out, details, err := prg.Eval(vars)
	if err != nil {
		logs.Error("file permission evaluation", logs.Details("details", details))
		return err
	}

	if out.Value().(bool) {
		return nil
	}

	return errors.Forbidden("permission denied")
}

func (h *PolicyHandler) assertIsAllowedToRead(ctx context.Context, sourceID string, filename string) error {
	user := auth.Get(ctx)
	if user != nil && user.Name == "admin" {
		return nil
	}

	source, err := h.next.GetSource(ctx, sourceID)
	if source == nil {
		return err
	}

	if source.PermissionOverrides != nil && len(source.PermissionOverrides.Read) > 0 {
		var rules []string
		for _, wr := range source.PermissionOverrides.Read {
			rules = append(rules, wr.Rule)
		}
		return h.assertPermissionIsGranted(ctx, rules...)
	}

	attrs, err := h.next.GetFileAttributes(ctx, sourceID, filename, AttrPermissions)
	if err != nil {
		return err
	}

	attrsHolder := HoldAttributes(attrs)
	perms, found, err := attrsHolder.GetPermissions()
	if err != nil {
		return err
	}

	if !found {
		return errors.Forbidden("access to this resources is forbidden")
	}

	var rules []string
	for _, perm := range perms.Read {
		rules = append(rules, perm.Rule)
	}

	return h.assertPermissionIsGranted(ctx, rules...)
}

func (h *PolicyHandler) assertIsAllowedToWrite(ctx context.Context, sourceID string, filename string) error {
	user := auth.Get(ctx)
	if user != nil && user.Name == "admin" {
		return nil
	}

	source, err := h.next.GetSource(ctx, sourceID)
	if err != nil {
		logs.Error("source not found", logs.Err(err))
		return err
	}

	if source.PermissionOverrides != nil && len(source.PermissionOverrides.Write) > 0 {
		var rules []string
		for _, wr := range source.PermissionOverrides.Write {
			rules = append(rules, wr.Rule)
		}
		return h.assertPermissionIsGranted(ctx, rules...)
	}

	attrs, err := h.next.GetFileAttributes(ctx, sourceID, filename, AttrPermissions)
	if err != nil {
		return err
	}

	attrsHolder := HoldAttributes(attrs)
	perms, found, err := attrsHolder.GetPermissions()
	if err != nil {
		return err
	}

	if !found {
		return errors.Forbidden("access to this resources is forbidden")
	}

	var rules []string
	for _, perm := range perms.Write {
		rules = append(rules, perm.Rule)
	}

	return h.assertPermissionIsGranted(ctx, rules...)
}

func (h *PolicyHandler) assertIsAllowedToChmod(ctx context.Context, sourceID string, filename string) error {
	user := auth.Get(ctx)
	if user != nil && user.Name == "admin" {
		return nil
	}

	source, err := h.next.GetSource(ctx, sourceID)
	if source == nil {
		return err
	}

	if source.PermissionOverrides != nil && len(source.PermissionOverrides.Chmod) > 0 {
		var rules []string
		for _, wr := range source.PermissionOverrides.Chmod {
			rules = append(rules, wr.Rule)
		}
		return h.assertPermissionIsGranted(ctx, rules...)
	}

	attrs, err := h.next.GetFileAttributes(ctx, sourceID, filename, AttrPermissions)
	if err != nil {
		return err
	}

	attrsHolder := HoldAttributes(attrs)
	perms, found, err := attrsHolder.GetPermissions()
	if err != nil {
		return err
	}

	if !found {
		return errors.Forbidden("access to this resources is forbidden")
	}

	var rules []string
	for _, perm := range perms.Chmod {
		rules = append(rules, perm.Rule)
	}

	return h.assertPermissionIsGranted(ctx, rules...)
}

func (h *PolicyHandler) assertAllowedToChmodSource(ctx context.Context, source *Source) error {
	user := auth.Get(ctx)
	if user == nil {
		return errors.Forbidden("")
	}

	if user.Name == "admin" {
		return nil
	}

	sourceChain := []string{source.Id}
	sourceType := source.Type
	var refSourceID string

	for sourceType == SourceType_Reference {
		u, err := url.Parse(source.Uri)
		if err != nil {
			return errors.Internal("could not resolve source uri", errors.Details{Key: "uri", Value: err})
		}

		if u.Scheme != "ref" {
			return errors.Internal("unexpected source scheme")
		}

		refSourceID = u.Host
		refSource, err := h.next.GetSource(ctx, refSourceID)
		if err != nil {
			return err
		}

		if refSource.PermissionOverrides != nil && len(refSource.PermissionOverrides.Chmod) > 0 {
			var rules []string
			for _, wr := range refSource.PermissionOverrides.Chmod {
				rules = append(rules, wr.Rule)
			}
			return h.assertPermissionIsGranted(ctx, rules...)
		}

		sourceType = refSource.Type

		for _, src := range sourceChain {
			if src == refSourceID {
				return errors.Internal("source cycle references")
			}
		}
		sourceChain = append(sourceChain, refSourceID)
		sourceType = source.Type
	}

	attrs, err := h.next.GetFileAttributes(ctx, refSourceID, "/", AttrPermissions)
	if err != nil {
		return err
	}

	attrsHolder := HoldAttributes(attrs)
	perms, found, err := attrsHolder.GetPermissions()
	if err != nil {
		return err
	}

	if !found {
		return errors.Forbidden("access to this resources is forbidden")
	}

	var rules []string
	for _, perm := range perms.Chmod {
		rules = append(rules, perm.Rule)
	}

	return h.assertPermissionIsGranted(ctx, rules...)
}

func (h *PolicyHandler) CreateSource(ctx context.Context, source *Source) error {
	clientApp := auth.App(ctx)
	if clientApp == nil || !clientApp.Sources.Create {
		return errors.Forbidden("application is not allowed to create sources")
	}

	err := h.assertAllowedToChmodSource(ctx, source)
	if err != nil {
		return err
	}
	return h.next.CreateSource(ctx, source)
}

func (h *PolicyHandler) ListSources(ctx context.Context) ([]*Source, error) {
	clientApp := auth.App(ctx)
	if clientApp == nil || !clientApp.Sources.Delete {
		return nil, errors.Forbidden("application is not allowed to list sources")
	}

	sources, err := h.next.ListSources(ctx)
	if err != nil {
		return nil, err
	}

	var allowedSources []*Source
	for _, source := range sources {
		err = h.assertIsAllowedToRead(ctx, source.Id, "/")
		if err != nil {
			continue
		}
		allowedSources = append(allowedSources, source)
	}
	return allowedSources, nil
}

func (h *PolicyHandler) GetSource(ctx context.Context, sourceID string) (*Source, error) {
	clientApp := auth.App(ctx)
	if clientApp == nil || !clientApp.Sources.Delete {
		return nil, errors.Forbidden("application is not allowed to list sources")
	}

	err := h.assertIsAllowedToRead(ctx, sourceID, "/")
	if err != nil {
		return nil, err
	}
	return h.next.GetSource(ctx, sourceID)
}

func (h *PolicyHandler) DeleteSource(ctx context.Context, sourceID string) error {
	user := auth.Get(ctx)
	if user == nil {
		return errors.Forbidden("context missing user")
	}

	clientApp := auth.App(ctx)
	if clientApp == nil || !clientApp.Sources.Delete {
		return errors.Forbidden("application is not allowed to delete sources")
	}

	source, err := h.next.GetSource(ctx, sourceID)
	if err != nil {
		return err
	}

	if user.Name != "admin" {
		if source.CreatedBy != user.Name {
			return errors.Forbidden("context missing user")
		}
	}
	return h.next.DeleteSource(ctx, sourceID)
}

func (h *PolicyHandler) CreateDir(ctx context.Context, sourceID string, filename string) error {
	err := h.assertIsAllowedToWrite(ctx, sourceID, path.Dir(filename))
	if err != nil {
		return err
	}
	return h.next.CreateDir(ctx, sourceID, filename)
}

func (h *PolicyHandler) WriteFileContent(ctx context.Context, sourceID string, filename string, content io.Reader, size int64, opts WriteOptions) error {
	err := h.assertIsAllowedToWrite(ctx, sourceID, path.Dir(filename))
	if err != nil {
		return err
	}
	err = h.next.WriteFileContent(ctx, sourceID, filename, content, size, opts)
	return err
}

func (h *PolicyHandler) ListDir(ctx context.Context, sourceID string, dirname string, opts ListDirOptions) (*DirContent, error) {
	err := h.assertIsAllowedToRead(ctx, sourceID, dirname)
	if err != nil {
		return nil, err
	}
	return h.next.ListDir(ctx, sourceID, dirname, opts)
}

func (h *PolicyHandler) ReadFileContent(ctx context.Context, sourceID string, filename string, opts ReadOptions) (io.ReadCloser, int64, error) {
	err := h.assertIsAllowedToRead(ctx, sourceID, filename)
	if err != nil {
		return nil, 0, err
	}
	return h.next.ReadFileContent(ctx, sourceID, filename, opts)
}

func (h *PolicyHandler) GetFileInfo(ctx context.Context, sourceID string, filename string, opts GetFileOptions) (*File, error) {
	err := h.assertIsAllowedToWrite(ctx, sourceID, path.Dir(filename))
	if err != nil {
		return nil, err
	}
	return h.next.GetFileInfo(ctx, sourceID, filename, opts)
}

func (h *PolicyHandler) DeleteFile(ctx context.Context, sourceID string, filename string, opts DeleteFileOptions) error {
	err := h.assertIsAllowedToWrite(ctx, sourceID, filename)
	if err != nil {
		return err
	}

	return h.next.DeleteFile(ctx, sourceID, filename, opts)
}

func (h *PolicyHandler) SetFileAttributes(ctx context.Context, sourceID string, filename string, attrs Attributes) error {
	err := h.assertIsAllowedToWrite(ctx, sourceID, filename)
	if err != nil {
		return err
	}
	return h.next.SetFileAttributes(ctx, sourceID, filename, attrs)
}

func (h *PolicyHandler) GetFileAttributes(ctx context.Context, sourceID string, filename string, name ...string) (Attributes, error) {
	err := h.assertIsAllowedToRead(ctx, sourceID, filename)
	if err != nil {
		return nil, err
	}
	return h.next.GetFileAttributes(ctx, sourceID, filename, name...)
}

func (h *PolicyHandler) RenameFile(ctx context.Context, sourceID string, filename string, newName string) error {
	err := h.assertIsAllowedToRead(ctx, sourceID, filename)
	if err != nil {
		return err
	}

	err = h.assertIsAllowedToWrite(ctx, sourceID, path.Dir(filename))
	if err != nil {
		return err
	}

	return h.next.RenameFile(ctx, sourceID, filename, newName)
}

func (h *PolicyHandler) MoveFile(ctx context.Context, sourceID string, filename string, dirname string) error {
	err := h.assertIsAllowedToWrite(ctx, sourceID, filename)
	if err != nil {
		return err
	}

	err = h.assertIsAllowedToWrite(ctx, sourceID, dirname)
	if err != nil {
		return err
	}

	return h.next.MoveFile(ctx, filename, sourceID, dirname)
}

func (h *PolicyHandler) CopyFile(ctx context.Context, sourceID string, filename string, dirname string) error {
	err := h.assertIsAllowedToRead(ctx, sourceID, filename)
	if err != nil {
		return err
	}

	err = h.assertIsAllowedToWrite(ctx, sourceID, dirname)
	if err != nil {
		return err
	}

	return h.next.CopyFile(ctx, sourceID, filename, dirname)
}

func (h *PolicyHandler) OpenMultipartSession(ctx context.Context, sourceID string, filename string, info MultipartSessionInfo) (string, error) {
	err := h.assertIsAllowedToWrite(ctx, sourceID, path.Dir(filename))
	if err != nil {
		return "", err
	}
	return h.next.OpenMultipartSession(ctx, sourceID, filename, info)
}

func (h *PolicyHandler) WriteFilePart(ctx context.Context, sessionID string, content io.Reader, size int64, info ContentPartInfo) (int64, error) {
	panic("implement me")
}

func (h *PolicyHandler) CloseMultipartSession(_ context.Context, _ string) error {
	panic("implement me")
}
