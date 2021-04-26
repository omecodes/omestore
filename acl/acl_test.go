package acl

import (
	"context"
	"database/sql"
	"github.com/omecodes/bome"
	pb "github.com/omecodes/store/gen/go/proto"
	"os"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	namespaceStore NamespaceConfigStore
	tupleStore     TupleStore

	namespaceDbConn *sql.DB
	tupleDBConn     *sql.DB

	err    error

	/*
		Doc namespace that describe relations for documents.
		parent: specifies parent of a document
		 owner: specifies owner of a document
		editor: specifies editor of a document. And tells that an owner is also a editor
		viewer: specifies viewer of a document. And tells that an editor and reader of document parent are readers too
	*/
	docNamespace = &pb.NamespaceConfig{
		Sid:       1,
		Namespace: "doc",
		Relations: map[string]*pb.RelationDefinition{
			"parent": {
				Name: "parent",
				SubjectSetRewrite: []*pb.SubjectSetDefinition{
					{
						Type: pb.SubjectSetType_This,
					},
				},
			},
			"owner": {
				Name: "owner",
				SubjectSetRewrite: []*pb.SubjectSetDefinition{
					{
						Type: pb.SubjectSetType_This,
					},
				},
			},
			"editor": {
				Name: "editor",
				SubjectSetRewrite: []*pb.SubjectSetDefinition{
					{
						Type: pb.SubjectSetType_This,
					},
					{
						Type:  pb.SubjectSetType_Computed,
						Value: "owner",
					},
				},
			},
			"viewer": {
				Name: "viewer",
				SubjectSetRewrite: []*pb.SubjectSetDefinition{
					{
						Type: pb.SubjectSetType_This,
					},
					{
						Type:  pb.SubjectSetType_Computed,
						Value: "editor",
					},
					{
						Type:  pb.SubjectSetType_FromTuple,
						Value: `{"object_relation":  "parent", "subject_relation":  "viewer"}`,
					},
				},
			},
		},
	}


	/*
		Situation: representing a doc tree /documents/
		/ (d0)
		|__ documents(d1)
		|	|__ description.pdf(d11)
		|	|__ course.pdf(d12)
		|__ images(d2)
		|	|__ pic.png (d21)
		|__ video(d3)
			|__ video.mp4 (d31)

		admin is owner of d0
		user1 is editor of d1
		user2 is editor of d2
		user3 is editor of d3

		Relations
	 */

	dataACL = []*pb.ACL {
		// Files parent relations
		{
			Subject:   "doc:d0",
			Relation: "parent",
			Object:  "doc:d1",
		},
		{
			Subject:   "doc:d0",
			Relation: "parent",
			Object:  "doc:d11",
		},
		{
			Subject:   "doc:d0",
			Relation: "parent",
			Object:  "doc:d12",
		},
		{
			Subject:   "doc:d0",
			Relation: "parent",
			Object:  "doc:d2",
		},
		{
			Subject:   "doc:d0",
			Relation: "parent",
			Object:  "doc:d21",
		},
		{
			Subject:   "doc:d0",
			Relation: "parent",
			Object:  "doc:d3",
		},
		{
			Subject:   "doc:d0",
			Relation: "parent",
			Object:  "doc:d31",
		},
		{
			Subject:   "doc:d1",
			Relation: "parent",
			Object:  "doc:d11",
		},
		{
			Subject:   "doc:d1",
			Relation: "parent",
			Object:  "doc:d12",
		},
		{
			Subject:   "doc:d2",
			Relation: "parent",
			Object:  "doc:d21",
		},
		{
			Subject:   "doc:d3",
			Relation: "parent",
			Object:  "doc:d31",
		},
		// Users relation with files
		{
			Subject:   "admin",
			Relation: "owner",
			Object:  "doc:d0",
		},
		{
			Subject:   "user1",
			Relation: "editor",
			Object:  "doc:d1",
		},
		{
			Subject:   "user2",
			Relation: "editor",
			Object:  "doc:d2",
		},
	}
)

func initDBs() {
	initNamespaceDB()
	initRelationDB()
}

func initNamespaceDB() {
	if namespaceDbConn == nil {
		dbFilename := "namespace.db"
		_ = os.Remove(dbFilename)

		namespaceDbConn, err = sql.Open("sqlite3", dbFilename)
		So(err, ShouldBeNil)

		namespaceStore, err = NewNamespaceSQLStore(namespaceDbConn, bome.SQLite3, "acl_")
		So(err, ShouldBeNil)
	}
}

func initRelationDB() {
	if tupleDBConn == nil {
		dbFilename := "tuples.db"
		_ = os.Remove(dbFilename)

		tupleDBConn, err = sql.Open("sqlite3", dbFilename)
		So(err, ShouldBeNil)

		tupleStore, err = NewTupleSQLStore(tupleDBConn, bome.SQLite3, "acl")
		So(err, ShouldBeNil)
	}
}

func closeTupleDBConn() {
	if tupleDBConn != nil {
		_ = tupleDBConn.Close()
		tupleDBConn = nil
	}
}

func closeNamespaceDBConn() {
	if namespaceDbConn != nil {
		_ = namespaceDbConn.Close()
		namespaceDbConn = nil
	}
}

func closeDBs() {
	closeTupleDBConn()
	closeNamespaceDBConn()
}

func fullManagerContext() context.Context {
	ctx := context.WithValue(context.Background(), ctxTupleStore{}, tupleStore)
	ctx = context.WithValue(ctx, ctxNamespaceConfigStore{}, namespaceStore)
	return ctx
}

func managerContextWithoutStore() context.Context {
	ctx := context.WithValue(context.Background(), ctxTupleStore{}, tupleStore)
	return ctx
}

func managerContextWithoutNamespace() context.Context {
	ctx := context.WithValue(context.Background(), ctxNamespaceConfigStore{}, namespaceStore)
	return ctx
}