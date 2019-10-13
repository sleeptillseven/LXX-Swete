package types

import "encoding/xml"

type Book struct {
	XMLName   xml.Name `xml:"TEI"`
	Chardata  string   `xml:",chardata"`
	Xmlns     string   `xml:"xmlns,attr"`
	TeiHeader struct {
		Text     string `xml:",chardata"`
		Lang     string `xml:"lang,attr"`
		FileDesc struct {
			Text      string `xml:",chardata"`
			TitleStmt struct {
				Text  string `xml:",chardata"`
				Title struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
				} `xml:"title"`
				Author struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
				} `xml:"author"`
				Editor    string `xml:"editor"`
				Funder    string `xml:"funder"`
				Principal string `xml:"principal"`
				RespStmt  []struct {
					Text     string `xml:",chardata"`
					PersName struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
					} `xml:"persName"`
					Resp string `xml:"resp"`
				} `xml:"respStmt"`
			} `xml:"titleStmt"`
			PublicationStmt struct {
				Text      string `xml:",chardata"`
				Authority string `xml:"authority"`
				Idno      struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"idno"`
				Availability struct {
					Text string `xml:",chardata"`
					P    string `xml:"p"`
				} `xml:"availability"`
				Date      string `xml:"date"`
				Publisher string `xml:"publisher"`
				PubPlace  string `xml:"pubPlace"`
			} `xml:"publicationStmt"`
			SourceDesc struct {
				Text     string `xml:",chardata"`
				ListBibl struct {
					Text       string `xml:",chardata"`
					BiblStruct struct {
						Text   string `xml:",chardata"`
						Monogr struct {
							Text   string `xml:",chardata"`
							Title  string `xml:"title"`
							Editor struct {
								Text     string `xml:",chardata"`
								PersName struct {
									Text string `xml:",chardata"`
									Name string `xml:"name"`
								} `xml:"persName"`
							} `xml:"editor"`
							Author struct {
								Text string `xml:",chardata"`
								Ref  string `xml:"ref,attr"`
								Lang string `xml:"lang,attr"`
							} `xml:"author"`
							Imprint struct {
								Text      string `xml:",chardata"`
								Publisher string `xml:"publisher"`
								PubPlace  string `xml:"pubPlace"`
								Date      string `xml:"date"`
							} `xml:"imprint"`
							BiblScope struct {
								Text string `xml:",chardata"`
								Unit string `xml:"unit,attr"`
							} `xml:"biblScope"`
						} `xml:"monogr"`
						Ref struct {
							Text   string `xml:",chardata"`
							Target string `xml:"target,attr"`
						} `xml:"ref"`
					} `xml:"biblStruct"`
				} `xml:"listBibl"`
			} `xml:"sourceDesc"`
		} `xml:"fileDesc"`
		EncodingDesc struct {
			Text     string `xml:",chardata"`
			P        string `xml:"p"`
			RefsDecl struct {
				Text        string `xml:",chardata"`
				N           string `xml:"n,attr"`
				CRefPattern []struct {
					Text               string `xml:",chardata"`
					MatchPattern       string `xml:"matchPattern,attr"`
					N                  string `xml:"n,attr"`
					ReplacementPattern string `xml:"replacementPattern,attr"`
				} `xml:"cRefPattern"`
			} `xml:"refsDecl"`
		} `xml:"encodingDesc"`
		ProfileDesc struct {
			Text      string `xml:",chardata"`
			LangUsage struct {
				Text     string `xml:",chardata"`
				Language []struct {
					Text  string `xml:",chardata"`
					Ident string `xml:"ident,attr"`
				} `xml:"language"`
			} `xml:"langUsage"`
		} `xml:"profileDesc"`
		RevisionDesc struct {
			Text   string `xml:",chardata"`
			Change string `xml:"change"`
		} `xml:"revisionDesc"`
	} `xml:"teiHeader"`
	Text struct {
		Text string `xml:",chardata"`
		Body struct {
			Text string `xml:",chardata"`
			Div  struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Lang string `xml:"lang,attr"`
				N    string `xml:"n,attr"`
				Pb   struct {
					Text string `xml:",chardata"`
					N    string `xml:"n,attr"`
				} `xml:"pb"`
				Head string `xml:"head"`
				Div  []struct {
					Text    string `xml:",chardata"`
					Type    string `xml:"type,attr"`
					Subtype string `xml:"subtype,attr"`
					N       string `xml:"n,attr"`
					Note    []struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"note"`
					Div []struct {
						Text    string `xml:",chardata"`
						Type    string `xml:"type,attr"`
						Subtype string `xml:"subtype,attr"`
						N       string `xml:"n,attr"`
						P       struct {
							Text string `xml:",chardata"`
							Note []struct {
								Text string `xml:",chardata"`
								Type string `xml:"type,attr"`
								Note struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"note"`
							} `xml:"note"`
							Pb struct {
								Text string `xml:",chardata"`
								N    string `xml:"n,attr"`
							} `xml:"pb"`
							Del struct {
								Text   string `xml:",chardata"`
								Status string `xml:"status,attr"`
							} `xml:"del"`
							Lg struct {
								Text string `xml:",chardata"`
								L    []struct {
									Text string `xml:",chardata"`
									Rend string `xml:"rend,attr"`
								} `xml:"l"`
							} `xml:"lg"`
						} `xml:"p"`
					} `xml:"div"`
					Pb struct {
						Text string `xml:",chardata"`
						N    string `xml:"n,attr"`
					} `xml:"pb"`
				} `xml:"div"`
			} `xml:"div"`
		} `xml:"body"`
	} `xml:"text"`
}
