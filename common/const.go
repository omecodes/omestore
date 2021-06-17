package common

const (
	AccessInfoEncryptedSecret     = "encrypted_secret"
	AccessInfoSecretEncryptParams = "encrypted_secret_params"
)

const (
	AdminAuthFile         = "./admin.auth"
	CookiesKeyFilename    = "./cookies.key"
	ServiceAuthSecretFile = "./services.auth"
)

const (
	CACertificateFilename = "service-ca.crt"
	CAKeyFilename         = "service-ca.key"
)

const (
	HttpHeaderContentType              = "Content-Type"
	HttpHeaderContentLength            = "Content-Length"
	HttpHeaderAccept                   = "Accept"
	HttpHeaderAccessControlAllowOrigin = "Access-Control-Allow-Origin"
	HttpHeaderUserAuthorization        = "Authorization"
	HttpHeaderAppAuthorization         = "X-STORE-CLIENT-APP-AUTHENTICATION"
)

const (
	ContentTypeJSONStream = "application/stream+json"
	ContentTypeJSON       = "application/json"
	AllJSONContentTypes   = ContentTypeJSONStream + "," + ContentTypeJSON
)

const (
	ApiDefaultLocation     = "/api"
	ApiObjectsRoutePrefix  = "/api/objects"
	ApiFilesRoutePrefix    = "/api/files"
	ApiAuthRoutePrefix     = "/api/auth"
	ApiAccountsRoutePrefix = "/api/accounts"
	ApiSettingsRoutePrefix = "/api/settings"
	ApiACLRoutePrefix      = "/api/acl"

	ApiQueryParamOffset = "offset"
	ApiParamAt          = "at"
	ApiParamName        = "name"
	ApiParamQuery       = "q"
	ApiParamUsername    = "username"
	ApiParamPassword    = "password"
	ApiQueryParamPath   = "path"
	ApiParamHeader      = "header"
	ApiParamContinueURL = "continue"

	ApiRouteVarId         = "{id}"
	ApiRouteVarKey        = "{key}"
	ApiRouteVarSource     = "{source}"
	ApiRouteVarName       = "{name}"
	ApiRouteVarCollection = "{collection}"

	ApiRouteVarIdName         = "id"
	ApiRouteVarKeyName        = "key"
	ApiRouteVarSourceName     = "source"
	ApiRouteVarNameName       = "name"
	ApiRouteVarCollectionName = "collection"

	// API Routes
	//

	ApiSetSettingsRoute = "/settings"
	ApiGetSettingsRoute = "/settings"

	ApiLoginRoute = "/login"

	ApiGetAccountRoute    = "/accounts/{id}"
	ApiCreateAccountRoute = "/accounts/{id}"
	ApiFindAccountRoute   = "/accounts/{id}"

	ApiSaveAuthProviderRoute   = "/auth/providers"
	ApiGetAuthProviderRoute    = "/auth/providers/{id}"
	ApiDeleteAuthProviderRoute = "/auth/providers/{id}"
	ApiListAuthProvidersRoute  = "/auth/providers"

	ApiSaveClientAppRoute    = "/auth/apps"
	ApiListClientAppsRoute   = "/auth/apps"
	ApiGetClientAppRoute     = "/auth/apps/{id}"
	ApiDeleteClientAppRoute  = "/auth/apps/{id}"
	ApiCreateAppSessionRoute = "/auth/sessions/client-app"
	ApiCreateUserRoute       = "/auth/users"
	ApiSearchUsersRoute      = "/auth/users"

	ApiCreateACLNamespaceConfigRoute         = "/acl/namespaces"
	ApiGetACLNamespaceConfigRoute            = "/acl/namespaces/{id}"
	ApiDeleteACLNamespaceConfigRoute         = "/acl/namespaces/{id}"
	ApiSaveACLRelationTupleRoute             = "/acl/relations"
	ApiDeleteACLRelationTupleRoute           = "/acl/relations/{id}"
	ApiCheckACLRelationTupleRoute            = "/acl/relations"
	ApiGetACLRelationTupleSubjectsNamesRoute = "/acl/relations/subjects"
	ApiGetACLRelationTupleObjectsNamesRoute  = "/acl/relations/objects"

	ApiACLNamespaceRoutePrefix = "/acl/namespaces"
	ApiACLRelationRoutePrefix  = "/acl/relations"

	ApiCreateCollectionRoute = "/objects/collections"
	ApiListCollectionRoute   = "/objects/collections"
	ApiGetCollectionRoute    = "/objects/collections/{id}"
	ApiDeleteCollectionRoute = "/objects/collections/{id}"
	ApiPutObjectRoute        = "/objects/data/{collection}"
	ApiPatchObjectRoute      = "/objects/data/{collection}/{id}"
	ApiMoveObjectRoute       = "/objects/data/{collection}/{id}"
	ApiGetObjectRoute        = "/objects/data/{collection}/{id}"
	ApiDeleteObjectRoute     = "/objects/data/{collection}/{id}"
	ApiListObjectsRoute      = "/objects/data/{collection}"
	ApiSearchObjectsRoute    = "/objects/data/{collection}"

	ApiCreateFileAccess          = "/files/accesses"
	ApiListFileAccesses          = "/files/accesses"
	ApiGetFileAccess             = "/files/accesses/{id}"
	ApiDeleteFileAccess          = "/files/accesses/{id}"
	ApiFileTreeRoutePrefix       = "/files/tree"
	ApiFileAttributesRoutePrefix = "/files/attributes"
	ApiFileDataRoutePrefix       = "/files/data"
)
