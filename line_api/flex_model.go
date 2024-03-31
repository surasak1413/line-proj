package line_api

// FlexContents represents the contents of a flex message
type FlexContents struct {
	Type     string        `json:"type"` // Type of the contents, e.g., "bubble"
	Styles   *FlexStyles   `json:"styles,omitempty"`
	Body     FlexBody      `json:"body"`
	Footer   *FlexFooter   `json:"footer,omitempty"`
	Hero     *FlexHero     `json:"hero,omitempty"`
	Header   *FlexHeader   `json:"header,omitempty"`
	Blocks   []FlexBlock   `json:"blocks,omitempty"`
	Contents []FlexContent `json:"contents,omitempty"`
}

// FlexStyles represents styles for a flex message
type FlexStyles struct {
	Header FlexHeaderStyles `json:"header,omitempty"`
	Footer FlexFooterStyles `json:"footer,omitempty"`
	Body   FlexBodyStyles   `json:"body,omitempty"`
}

// FlexHeader represents the header of a flex message
type FlexHeader struct {
	Type     string      `json:"type"` // Type of the header, e.g., "box"
	Layout   string      `json:"layout,omitempty"`
	Contents []FlexBlock `json:"contents,omitempty"`
}

// FlexFooter represents the footer of a flex message
type FlexFooter struct {
	Type     string      `json:"type"` // Type of the footer, e.g., "box"
	Layout   string      `json:"layout,omitempty"`
	Contents []FlexBlock `json:"contents,omitempty"`
}

// FlexHero represents the hero component of a flex message
type FlexHero struct {
	Type        string     `json:"type"` // Type of the hero, e.g., "image"
	URL         string     `json:"url,omitempty"`
	Size        string     `json:"size,omitempty"`
	AspectRatio string     `json:"aspectRatio,omitempty"`
	AspectMode  string     `json:"aspectMode,omitempty"`
	Action      FlexAction `json:"action,omitempty"`
}

// FlexBody represents the body component of a flex message
type FlexBody struct {
	Type     string      `json:"type"` // Type of the body, e.g., "box"
	Layout   string      `json:"layout,omitempty"`
	Contents []FlexBlock `json:"contents,omitempty"`
}

// FlexBlock represents a generic block in a flex message
type FlexBlock struct {
	Type     string      `json:"type"` // Type of the block, e.g., "box"
	Layout   string      `json:"layout,omitempty"`
	Contents []FlexBlock `json:"contents,omitempty"`
	Text     string      `json:"text,omitempty"`
	Wrap     bool        `json:"wrap,omitempty"`
	Weight   string      `json:"weight,omitempty"`
	URL      string      `json:"url,omitempty"`
	Margin   string      `json:"margin,omitempty"`
	Spacing  string      `json:"spacing,omitempty"`
	Size     string      `json:"size,omitempty"`
	Color    string      `json:"color,omitempty"`
	Align    string      `json:"align,omitempty"`
	Flex     int         `json:"flex,omitempty"`
}

// FlexContent represents a content in a flex message
type FlexContent struct {
	Type     string      `json:"type"` // Type of the content, e.g., "text"
	Text     string      `json:"text,omitempty"`
	Contents []FlexBlock `json:"contents,omitempty"`
	Action   FlexAction  `json:"action,omitempty"`
}

// FlexAction represents an action in a flex message
type FlexAction struct {
	Type  string `json:"type"` // Type of the action, e.g., "message"
	Label string `json:"label,omitempty"`
	Text  string `json:"text,omitempty"`
	URI   string `json:"uri,omitempty"`
}

// FlexHeaderStyles represents styles for the header component
type FlexHeaderStyles struct {
	BackgroundColor string `json:"backgroundColor,omitempty"`
	Separator       bool   `json:"separator,omitempty"`
}

// FlexFooterStyles represents styles for the footer component
type FlexFooterStyles struct {
	BackgroundColor string `json:"backgroundColor,omitempty"`
	Separator       bool   `json:"separator,omitempty"`
}

// FlexBodyStyles represents styles for the body component
type FlexBodyStyles struct {
	BackgroundColor string `json:"backgroundColor,omitempty"`
}
