package sys_model

type UEditorConfig struct {
	State                   string   `json:"state"`
	ImageActionName         string   `json:"imageActionName"`
	ImageFieldName          string   `json:"imageFieldName"`
	ImageMaxSize            int      `json:"imageMaxSize"`
	ImageAllowFiles         []string `json:"imageAllowFiles"`
	ImageCompressEnable     bool     `json:"imageCompressEnable"`
	ImageCompressBorder     int      `json:"imageCompressBorder"`
	ImageInsertAlign        string   `json:"imageInsertAlign"`
	ImageUrlPrefix          string   `json:"imageUrlPrefix"`
	ScrawlActionName        string   `json:"scrawlActionName"`
	ScrawlFieldName         string   `json:"scrawlFieldName"`
	ScrawlMaxSize           int      `json:"scrawlMaxSize"`
	ScrawlUrlPrefix         string   `json:"scrawlUrlPrefix"`
	ScrawlInsertAlign       string   `json:"scrawlInsertAlign"`
	SnapScreenActionName    string   `json:"snapScreenActionName"`
	SnapScreenUrlPrefix     string   `json:"snapScreenUrlPrefix"`
	SnapScreenInsertAlign   string   `json:"snapScreenInsertAlign"`
	CatcherLocalDomain      []string `json:"catcherLocalDomain"`
	CatcherActionName       string   `json:"catcherActionName"`
	CatcherFieldName        string   `json:"catcherFieldName"`
	CatcherUrlPrefix        string   `json:"catcherUrlPrefix"`
	CatcherMaxSize          int      `json:"catcherMaxSize"`
	CatcherAllowFiles       []string `json:"catcherAllowFiles"`
	VideoActionName         string   `json:"videoActionName"`
	VideoFieldName          string   `json:"videoFieldName"`
	VideoUrlPrefix          string   `json:"videoUrlPrefix"`
	VideoMaxSize            int      `json:"videoMaxSize"`
	VideoAllowFiles         []string `json:"videoAllowFiles"`
	FileActionName          string   `json:"fileActionName"`
	FileFieldName           string   `json:"fileFieldName"`
	FileUrlPrefix           string   `json:"fileUrlPrefix"`
	FileMaxSize             int      `json:"fileMaxSize"`
	FileAllowFiles          []string `json:"fileAllowFiles"`
	ImageManagerActionName  string   `json:"imageManagerActionName"`
	ImageManagerListSize    int      `json:"imageManagerListSize"`
	ImageManagerUrlPrefix   string   `json:"imageManagerUrlPrefix"`
	ImageManagerInsertAlign string   `json:"imageManagerInsertAlign"`
	ImageManagerAllowFiles  []string `json:"imageManagerAllowFiles"`
	FileManagerActionName   string   `json:"fileManagerActionName"`
	FileManagerUrlPrefix    string   `json:"fileManagerUrlPrefix"`
	FileManagerListSize     int      `json:"fileManagerListSize"`
	FileManagerAllowFiles   []string `json:"fileManagerAllowFiles"`
	FormulaConfig           struct {
		ImageUrlTemplate string `json:"imageUrlTemplate"`
	} `json:"formulaConfig"`
}
