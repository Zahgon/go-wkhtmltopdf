package wkhtmltopdf

const opt = "--"

// A list of options that can be set from code to make it easier to see which options are available
type globalOptions struct {
	CookieJar         stringOption // Read and write cookies from and to the supplied cookie jar file
	Copies            uintOption   // Number of copies to print into the pdf file (default 1)
	Dpi               uintOption   // Change the dpi explicitly (this has no effect on X11 based systems)
	ExtendedHelp      boolOption   // Display more extensive help, detailing less common command switches
	Grayscale         boolOption   // PDF will be generated in grayscale
	Help              boolOption   // Display help
	HTMLDoc           boolOption   // Output program html help
	ImageDpi          uintOption   // When embedding images scale them down to this dpi (default 600)
	ImageQuality      uintOption   // When jpeg compressing images use this quality (default 94)
	License           boolOption   // Output license information and exit
	LogLevel          stringOption // Set log level to: none, error, warn or info (default info)
	LowQuality        boolOption   // Generates lower quality pdf/ps. Useful to shrink the result document space
	ManPage           boolOption   // Output program man page
	MarginBottom      uintOption   // Set the page bottom margin in centimeters
	MarginBottomUnit  stringOption // Set the page bottom margin with a unit
	MarginLeft        uintOption   // Set the page left margin (default 10mm)
	MarginLeftUnit    stringOption // Set the page left margin with a unit
	MarginRight       uintOption   // Set the page right margin (default 10mm)
	MarginRightUnit   stringOption // Set the page right margin with a unit
	MarginTop         uintOption   // Set the page top margin
	MarginTopUnit     stringOption // Set the page top margin with a unit
	NoCollate         boolOption   // Do not collate when printing multiple copies (default collate)
	NoPdfCompression  boolOption   // Do not use lossless compression on pdf objects
	Orientation       stringOption // Set orientation to Landscape or Portrait (default Portrait)
	PageHeight        uintOption   // Page height
	PageHeightUnit    stringOption // Page height  with a unit
	PageSize          stringOption // Set paper size to: A4, Letter, etc. (default A4)
	PageWidth         uintOption   // Page width
	PageWidthUnit     stringOption // Page width with a unit
	Quiet             boolOption   // Be less verbose
	ReadArgsFromStdin boolOption   // Read command line arguments from stdin
	Readme            boolOption   // Output program readme
	Title             stringOption // The title of the generated pdf file (The title of the first document is used if not specified)
	Version           boolOption   // Output version information and exit
}

func (gopt *globalOptions) Args() []string { _ = "STUB: not implemented"; return nil }

type outlineOptions struct {
	DumpDefaultTocXsl boolOption   // Dump the default TOC xsl style sheet to stdout
	DumpOutline       stringOption // Dump the outline to a file
	NoOutline         boolOption   // Do not put an outline into the pdf
	OutlineDepth      uintOption   // Set the depth of the outline (default 4)
}

func (oopt *outlineOptions) Args() []string { _ = "STUB: not implemented"; return nil }

type pageOptions struct {
	Allow                     sliceOption  // Allow the file or files from the specified folder to be loaded (repeatable)
	BypassProxyFor            sliceOption  // Bypass proxy for host
	CacheDir                  stringOption // Web cache directory
	CheckboxCheckedSvg        stringOption // Use this SVG file when rendering checked checkboxes
	CheckboxSvg               stringOption // Use this SVG file when rendering unchecked checkboxes
	Cookie                    mapOption    // Set an additional cookie (repeatable), value should be url encoded
	CustomHeader              mapOption    // Set an additional HTTP header (repeatable)
	CustomHeaderPropagation   boolOption   // Add HTTP headers specified by --custom-header for each resource request
	DebugJavascript           boolOption   // Show javascript debugging output
	DefaultHeader             boolOption   // Add a default header, with the name of the page to the left, and the page number to the right, this is short for: --header-left='[webpage]' --header-right='[page]/[toPage]' --top 2cm --header-line
	DisableExternalLinks      boolOption   // Do not make links to remote web pages
	DisableInternalLinks      boolOption   // Do not make local links
	DisableJavascript         boolOption   // Do not allow web pages to run javascript
	DisableLocalFileAccess    boolOption   // Do not allowed conversion of a local file to read in other local files, unless explicitly allowed with --allow
	DisableSmartShrinking     boolOption   // Disable the intelligent shrinking strategy used by WebKit that makes the pixel/dpi ratio none constant
	EnableForms               boolOption   // Turn HTML form fields into pdf form fields
	EnableLocalFileAccess     boolOption   // Allowed conversion of a local file to read in other local files
	EnablePlugins             boolOption   // Enable installed plugins (plugins will likely not work)
	EnableTocBackLinks        boolOption   // Link from section header to toc
	Encoding                  stringOption // Set the default text encoding, for input
	ExcludeFromOutline        boolOption   // Do not include the page in the table of contents and outlines
	JavascriptDelay           uintOption   // Wait some milliseconds for javascript finish (default 200)
	KeepRelativeLinks         boolOption   // Keep relative external links as relative external links
	LoadErrorHandling         stringOption // Specify how to handle pages that fail to load: abort, ignore or skip (default abort)
	LoadMediaErrorHandling    stringOption // Specify how to handle media files that fail to load: abort, ignore or skip (default ignore)
	MinimumFontSize           uintOption   // Minimum font size
	NoBackground              boolOption   // Do not print background
	NoCustomHeaderPropagation boolOption   // Do not add HTTP headers specified by --custom-header for each resource request
	NoImages                  boolOption   // Do not load or print images
	NoStopSlowScripts         boolOption   // Do not Stop slow running javascripts
	PageOffset                uintOption   // Set the starting page number (default 0)
	Password                  stringOption // HTTP Authentication password
	Post                      mapOption    // Add an additional post field (repeatable)
	PostFile                  mapOption    // Post an additional file (repeatable)
	PrintMediaType            boolOption   // Use print media-type instead of screen
	Proxy                     stringOption // Use a proxy
	ProxyHostnameLookup       boolOption   // Use the proxy for resolving hostnames
	RadiobuttonCheckedSvg     stringOption // Use this SVG file when rendering checked radiobuttons
	RadiobuttonSvg            stringOption // Use this SVG file when rendering unchecked radiobuttons
	RunScript                 sliceOption  // Run this additional javascript after the page is done loading (repeatable)
	SslCrtPath                stringOption // Path to the ssl client cert public key in OpenSSL PEM format, optionally followed by intermediate ca and trusted certs
	SslKeyPassword            stringOption // Password to ssl client cert private key
	SslKeyPath                stringOption // Path to ssl client cert private key in OpenSSL PEM format
	Username                  stringOption // HTTP Authentication username
	UserStyleSheet            stringOption // Specify a user style sheet, to load with every page
	ViewportSize              stringOption // Set viewport size if you have custom scrollbars or css attribute overflow to emulate window size
	WindowStatus              stringOption // Wait until window.status is equal to this string before rendering page
	Zoom                      floatOption  // Use this zoom factor (default 1)
}

func (popt *pageOptions) Args() []string { _ = "STUB: not implemented"; return nil }

type headerAndFooterOptions struct {
	FooterCenter   stringOption // Centered footer text
	FooterFontName stringOption // Set footer font name (default Arial)
	FooterFontSize uintOption   // Set footer font size (default 12)
	FooterHTML     stringOption // Adds a html footer
	FooterLeft     stringOption // Left aligned footer text
	FooterLine     boolOption   // Display line above the footer
	FooterRight    stringOption // Right aligned footer text
	FooterSpacing  floatOption  // Spacing between footer and content in mm (default 0)
	HeaderCenter   stringOption // Centered header text
	HeaderFontName stringOption // Set header font name (default Arial)
	HeaderFontSize uintOption   // Set header font size (default 12)
	HeaderHTML     stringOption // Adds a html header
	HeaderLeft     stringOption // Left aligned header text
	HeaderLine     boolOption   // Display line below the header
	HeaderRight    stringOption // Right aligned header text
	HeaderSpacing  floatOption  // Spacing between header and content in mm (default 0)
	Replace        mapOption    // Replace [name] with value in header and footer (repeatable)
}

func (hopt *headerAndFooterOptions) Args() []string { _ = "STUB: not implemented"; return nil }

type tocOptions struct {
	DisableDottedLines  boolOption   // Do not use dotted lines in the toc
	DisableTocLinks     boolOption   // Do not link from toc to sections
	TocHeaderText       stringOption // The header text of the toc (default Table of Contents)
	TocLevelIndentation uintOption   // For each level of headings in the toc indent by this length (default 1em)
	TocTextSizeShrink   floatOption  // For each level of headings in the toc the font is scaled by this factor
	XslStyleSheet       stringOption // Use the supplied xsl style sheet for printing the table of content
}

func (topt *tocOptions) Args() []string { _ = "STUB: not implemented"; return nil }

type argParser interface {
	Parse() []string //  Used in the cmd call
}

type stringOption struct {
	option string
	value  string
}

func (so stringOption) Parse() []string { _ = "STUB: not implemented"; return nil }

func (so *stringOption) Set(value string) { _ = "STUB: not implemented"; return }

func (so *stringOption) Unset() { _ = "STUB: not implemented"; return }

type sliceOption struct {
	option string
	value  []string
}

func (so sliceOption) Parse() []string { _ = "STUB: not implemented"; return nil }

func (so *sliceOption) Set(value string) { _ = "STUB: not implemented"; return }

func (so *sliceOption) Unset() { _ = "STUB: not implemented"; return }

type mapOption struct {
	option string
	value  map[string]string
}

func (mo mapOption) Parse() []string { _ = "STUB: not implemented"; return nil }

func (mo *mapOption) Set(key, value string) { _ = "STUB: not implemented"; return }

func (mo *mapOption) Unset() { _ = "STUB: not implemented"; return }

type uintOption struct {
	option string
	value  uint
	isSet  bool
}

func (io uintOption) Parse() []string { _ = "STUB: not implemented"; return nil }

func (io *uintOption) Set(value uint) { _ = "STUB: not implemented"; return }

func (io *uintOption) Unset() { _ = "STUB: not implemented"; return }

type floatOption struct {
	option string
	value  float64
	isSet  bool
}

func (fo floatOption) Parse() []string { _ = "STUB: not implemented"; return nil }

func (fo *floatOption) Set(value float64) { _ = "STUB: not implemented"; return }

func (fo *floatOption) Unset() { _ = "STUB: not implemented"; return }

type boolOption struct {
	option string
	value  bool
}

func (bo boolOption) Parse() []string { _ = "STUB: not implemented"; return nil }

func (bo *boolOption) Set(value bool) { _ = "STUB: not implemented"; return }

func (bo *boolOption) Unset() { _ = "STUB: not implemented"; return }

func newGlobalOptions() globalOptions { _ = "STUB: not implemented"; return *new(globalOptions) }

func newOutlineOptions() outlineOptions { _ = "STUB: not implemented"; return *new(outlineOptions) }

func newPageOptions() pageOptions { _ = "STUB: not implemented"; return *new(pageOptions) }

func newHeaderAndFooterOptions() headerAndFooterOptions {
	_ = "STUB: not implemented"
	return *new(headerAndFooterOptions)
}

func newTocOptions() tocOptions { _ = "STUB: not implemented"; return *new(tocOptions) }

func optsToArgs(opts interface{}) []string { _ = "STUB: not implemented"; return nil }

// Constants for orientation modes
const (
	OrientationLandscape = "Landscape" // Landscape mode
	OrientationPortrait  = "Portrait"  // Portrait mode
)

// Constants for page sizes
const (
	PageSizeA0        = "A0"        //	841 x 1189 mm
	PageSizeA1        = "A1"        //	594 x 841 mm
	PageSizeA2        = "A2"        //	420 x 594 mm
	PageSizeA3        = "A3"        //	297 x 420 mm
	PageSizeA4        = "A4"        //	210 x 297 mm, 8.26
	PageSizeA5        = "A5"        //	148 x 210 mm
	PageSizeA6        = "A6"        //	105 x 148 mm
	PageSizeA7        = "A7"        //	74 x 105 mm
	PageSizeA8        = "A8"        //	52 x 74 mm
	PageSizeA9        = "A9"        //	37 x 52 mm
	PageSizeB0        = "B0"        //	1000 x 1414 mm
	PageSizeB1        = "B1"        //	707 x 1000 mm
	PageSizeB10       = "B10"       //	31 x 44 mm
	PageSizeB2        = "B2"        //	500 x 707 mm
	PageSizeB3        = "B3"        //	353 x 500 mm
	PageSizeB4        = "B4"        //	250 x 353 mm
	PageSizeB5        = "B5"        //	176 x 250 mm, 6.93
	PageSizeB6        = "B6"        //	125 x 176 mm
	PageSizeB7        = "B7"        //	88 x 125 mm
	PageSizeB8        = "B8"        //	62 x 88 mm
	PageSizeB9        = "B9"        //	33 x 62 mm
	PageSizeC5E       = "C5E"       //	163 x 229 mm
	PageSizeComm10E   = "Comm10E"   //	105 x 241 mm, U.S. Common 10 Envelope
	PageSizeCustom    = "Custom"    //	Unknown, or a user defined size.
	PageSizeDLE       = "DLE"       //	110 x 220 mm
	PageSizeExecutive = "Executive" //	7.5 x 10 inches, 190.5 x 254 mm
	PageSizeFolio     = "Folio"     //	210 x 330 mm
	PageSizeLedger    = "Ledger"    //	431.8 x 279.4 mm
	PageSizeLegal     = "Legal"     //	8.5 x 14 inches, 215.9 x 355.6 mm
	PageSizeLetter    = "Letter"    //	8.5 x 11 inches, 215.9 x 279.4 mm
	PageSizeTabloid   = "Tabloid"   //	279.4 x 431.8 mm
)
