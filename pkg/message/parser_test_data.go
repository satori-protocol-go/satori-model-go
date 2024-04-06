package message

type rawToElement struct {
	Raw       string
	TargetRaw string
	Elements  []MessageElement
}

func _getRawMessage() map[string][]rawToElement {
	raw_message := make(map[string][]rawToElement)
	raw_message["Basic"] = _getBasicRawMessage()
	raw_message["Layout"] = _getLayoutRawMessage()
	raw_message["Meta"] = _getMetaRawMessage()
	raw_message["Resource"] = _getResourceRawMessage()
	raw_message["Modifier"] = _getModifierRawMessage()
	raw_message["Examples"] = _getExamplesRawMessage()
	return raw_message
}

func _getExamplesRawMessage() []rawToElement {
	return []rawToElement{
		// Single message
		{
			Raw: `
				<message>
				<at id="1" name="test" />
				<at id="2" name="test2" />
				<br/>
				<p>test</p>
				<img src="https://example.com" />
				<audio src="https://example.com" />
				<p>
					<b>bold</b>
					<sup>sup</sup>
					<sub>sub</sub>
					<spl>spoiler</spl>
					<code>code</code>
					<u>underline</u>
					<s>strikethrough</s>
				</p>
				</message>`,
			TargetRaw: `<message><at id="1" name="test" /><at id="2" name="test2" /><br/><p>test</p><img src="https://example.com" /><audio src="https://example.com" /><p><b>bold</b><sup>sup</sup><sub>sub</sub><spl>spoiler</spl><code>code</code><u>underline</u><s>strikethrough</s></p></message>`,
			Elements: []MessageElement{
				&MessageElementMessage{
					ChildrenMessageElement: &ChildrenMessageElement{
						Children: []MessageElement{
							&MessageElementAt{
								Id:   "1",
								Name: "test",
							},
							&MessageElementAt{
								Id:   "2",
								Name: "test2",
							},
							&MessageElmentBr{},
							&MessageElmentP{
								ChildrenMessageElement: &ChildrenMessageElement{
									Children: []MessageElement{
										&MessageElementText{
											Content: "test",
										},
									},
								},
							},
							&MessageElementImg{
								Src: "https://example.com",
							},
							&MessageElementAudio{
								Src: "https://example.com",
							},
							&MessageElmentP{
								ChildrenMessageElement: &ChildrenMessageElement{
									/**
									<b>bold</b>
									<sup>sup</sup>
									<sub>sub</sub>
									<spl>spoiler</spl>
									<code>code</code>
									<u>underline</u>
									<s>strikethrough</s>
									*/
									Children: []MessageElement{
										&MessageElementStrong{
											ChildrenMessageElement: &ChildrenMessageElement{
												Children: []MessageElement{
													&MessageElementText{
														Content: "bold",
													},
												},
											},
										},
										&MessageElementSup{
											ChildrenMessageElement: &ChildrenMessageElement{
												Children: []MessageElement{
													&MessageElementText{
														Content: "sup",
													},
												},
											},
										},
										&MessageElementSub{
											ChildrenMessageElement: &ChildrenMessageElement{
												Children: []MessageElement{
													&MessageElementText{
														Content: "sub",
													},
												},
											},
										},
										&MessageElementSpl{
											ChildrenMessageElement: &ChildrenMessageElement{
												Children: []MessageElement{
													&MessageElementText{
														Content: "spoiler",
													},
												},
											},
										},
										&MessageElementCode{
											ChildrenMessageElement: &ChildrenMessageElement{
												Children: []MessageElement{
													&MessageElementText{
														Content: "code",
													},
												},
											},
										},
										&MessageElementIns{
											ChildrenMessageElement: &ChildrenMessageElement{
												Children: []MessageElement{
													&MessageElementText{
														Content: "underline",
													},
												},
											},
										},
										&MessageElementDel{
											ChildrenMessageElement: &ChildrenMessageElement{
												Children: []MessageElement{
													&MessageElementText{
														Content: "strikethrough",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{Raw: `<message>
<code>
  yarn add nitori
  yarn run test
  npm publish
</code>
</message>`,
			TargetRaw: "<message><code>yarn add nitori\n  yarn run test\n  npm publish</code></message>",
			Elements: []MessageElement{
				&MessageElementMessage{
					ChildrenMessageElement: &ChildrenMessageElement{
						Children: []MessageElement{
							&MessageElementCode{
								ChildrenMessageElement: &ChildrenMessageElement{
									Children: []MessageElement{
										&MessageElementText{
											Content: "yarn add nitori\n  yarn run test\n  npm publish",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func _getModifierRawMessage() []rawToElement {
	return []rawToElement{
		{
			Raw:       `<strong>test</strong>`,
			TargetRaw: `<b>test</b>`,
			Elements: []MessageElement{&MessageElementStrong{
				&ChildrenMessageElement{
					Children: []MessageElement{&MessageElementText{
						Content: "test",
					}},
				},
			}},
		},
		{
			Raw: `<b>test</b>`,
			Elements: []MessageElement{&MessageElementStrong{
				&ChildrenMessageElement{
					Children: []MessageElement{&MessageElementText{
						Content: "test",
					}},
				},
			}},
		},
		{
			Raw:       `<em>test</em>`,
			TargetRaw: `<i>test</i>`,
			Elements: []MessageElement{&MessageElementEm{
				&ChildrenMessageElement{
					Children: []MessageElement{&MessageElementText{
						Content: "test",
					}},
				},
			}},
		},
		{
			Raw: `<i>test</i>`,
			Elements: []MessageElement{&MessageElementEm{
				&ChildrenMessageElement{
					Children: []MessageElement{&MessageElementText{
						Content: "test",
					}},
				},
			}},
		},
		{
			Raw:       `<ins>test</ins>`,
			TargetRaw: `<u>test</u>`,
			Elements: []MessageElement{&MessageElementIns{
				&ChildrenMessageElement{
					Children: []MessageElement{&MessageElementText{
						Content: "test",
					}},
				},
			}},
		},
		{
			Raw: `<u>test</u>`,
			Elements: []MessageElement{&MessageElementIns{
				&ChildrenMessageElement{
					Children: []MessageElement{&MessageElementText{
						Content: "test",
					}},
				},
			}},
		},
		{
			Raw:       `<del>test</del>`,
			TargetRaw: `<s>test</s>`,
			Elements: []MessageElement{&MessageElementDel{
				&ChildrenMessageElement{
					Children: []MessageElement{&MessageElementText{
						Content: "test",
					}},
				},
			}},
		},
		{
			Raw: `<s>test</s>`,
			Elements: []MessageElement{&MessageElementDel{
				&ChildrenMessageElement{
					Children: []MessageElement{&MessageElementText{
						Content: "test",
					}},
				},
			}},
		},
		{
			Raw: `<spl>test</spl>`,
			Elements: []MessageElement{&MessageElementSpl{
				ChildrenMessageElement: &ChildrenMessageElement{
					Children: []MessageElement{&MessageElementText{
						Content: "test",
					}},
				},
			}},
		},
		{
			Raw: `<code>test</code>`,
			Elements: []MessageElement{&MessageElementCode{
				ChildrenMessageElement: &ChildrenMessageElement{
					Children: []MessageElement{&MessageElementText{
						Content: "test",
					}},
				},
			}},
		},
		{
			Raw: `<sup>test</sup>`,
			Elements: []MessageElement{&MessageElementSup{
				ChildrenMessageElement: &ChildrenMessageElement{
					Children: []MessageElement{&MessageElementText{
						Content: "test",
					}},
				},
			}},
		},
		{
			Raw: `<sub>test</sub>`,
			Elements: []MessageElement{&MessageElementSub{
				ChildrenMessageElement: &ChildrenMessageElement{
					Children: []MessageElement{&MessageElementText{
						Content: "test",
					}},
				},
			}},
		},
	}
}

func _getResourceRawMessage() []rawToElement {
	return []rawToElement{
		{
			Raw: `<img src="https://example.com" />`,
			Elements: []MessageElement{
				&MessageElementImg{
					Src: "https://example.com",
				},
			},
		},
		{
			Raw: `<img src="https://example.com" cache timeout="1000" />`,
			Elements: []MessageElement{
				&MessageElementImg{
					Src:     "https://example.com",
					Cache:   true,
					Timeout: "1000",
				}},
		},
		{
			Raw:       `<img src="https://example.com" width="300" height="200" />`,
			TargetRaw: `<img src="https://example.com" width=300 height=200 />`,
			Elements: []MessageElement{&MessageElementImg{
				Src:    "https://example.com",
				Width:  300,
				Height: 200,
			}},
		},
		{
			Raw: `<audio src="https://example.com" />`,
			Elements: []MessageElement{&MessageElementAudio{
				Src: "https://example.com",
			}},
		},
		{
			Raw: `<video src="https://example.com" />`,
			Elements: []MessageElement{&MessageElementVideo{
				Src: "https://example.com",
			}},
		},
		{
			Raw: `<file src="https://example.com" />`,
			Elements: []MessageElement{&MessageElementFile{
				Src: "https://example.com",
			}},
		},
	}
}

func _getMetaRawMessage() []rawToElement {
	return []rawToElement{
		{
			Raw:      `<quote />`,
			Elements: []MessageElement{&MessageElementQuote{}},
		},
		{
			Raw: `<quote>test</quote>`,
			Elements: []MessageElement{&MessageElementQuote{
				ChildrenMessageElement: &ChildrenMessageElement{
					Children: []MessageElement{
						&MessageElementText{
							Content: "test",
						},
					},
				},
			}},
		},
		{
			Raw:      `<author />`,
			Elements: []MessageElement{&MessageElementAuthor{}},
		},
		{
			Raw: `<author id="1" />`,
			Elements: []MessageElement{&MessageElementAuthor{
				Id: "1",
			}},
		},
		{
			Raw: `<author id="1" name="test" />`,
			Elements: []MessageElement{&MessageElementAuthor{
				Id:   "1",
				Name: "test",
			}},
		},
	}
}

func _getLayoutRawMessage() []rawToElement {
	return []rawToElement{
		{
			Raw: "<br/>",
			Elements: []MessageElement{
				&MessageElmentBr{},
			},
		},
		{
			Raw: `<p>test</p>`,
			Elements: []MessageElement{
				&MessageElmentP{
					ChildrenMessageElement: &ChildrenMessageElement{
						Children: []MessageElement{
							&MessageElementText{
								Content: "test",
							},
						},
					},
				},
			},
		},
		{
			Raw: `<p>testtest2</p>`,
			Elements: []MessageElement{
				&MessageElmentP{
					ChildrenMessageElement: &ChildrenMessageElement{
						Children: []MessageElement{
							&MessageElementText{
								Content: "testtest2",
							},
						},
					},
				},
			},
		},
		{
			Raw: `<p>test<br/></p>`,
			Elements: []MessageElement{
				&MessageElmentP{
					ChildrenMessageElement: &ChildrenMessageElement{
						Children: []MessageElement{
							&MessageElementText{
								Content: "test",
							},
							&MessageElmentBr{},
						},
					},
				},
			},
		},
		{
			Raw: `<message />`,
			Elements: []MessageElement{
				&MessageElementMessage{},
			},
		},
		{
			Raw: `<message id="1" />`,
			Elements: []MessageElement{
				&MessageElementMessage{
					Id: "1",
				},
			},
		},
		{
			Raw: `<message forward />`,
			Elements: []MessageElement{
				&MessageElementMessage{
					Forward: true,
				},
			},
		},
		{
			Raw: `<message id="1" forward />`,
			Elements: []MessageElement{
				&MessageElementMessage{
					Id:      "1",
					Forward: true,
				},
			},
		},
		{
			Raw:       `<message id="1" forward="false" />`,
			TargetRaw: `<message id="1" />`,
			Elements: []MessageElement{
				&MessageElementMessage{
					Id:      "1",
					Forward: false,
				},
			},
		},
		{
			Raw: `<message>test</message>`,
			Elements: []MessageElement{
				&MessageElementMessage{
					ChildrenMessageElement: &ChildrenMessageElement{
						Children: []MessageElement{
							&MessageElementText{
								Content: "test",
							},
						},
					},
				},
			},
		},
	}
}

func _getBasicRawMessage() []rawToElement {
	return []rawToElement{
		{
			Raw: "纯文本",
			Elements: []MessageElement{
				&MessageElementText{
					Content: "纯文本",
				},
			},
		},
		{
			Raw: `<at id="1" name="test" />`,
			Elements: []MessageElement{
				&MessageElementAt{
					Id:   "1",
					Name: "test",
				},
			},
		},
		{
			Raw: `<at role="test" />`,
			Elements: []MessageElement{
				&MessageElementAt{
					Role: "test",
				},
			},
		},
		{
			Raw: `<at type="test" />`,
			Elements: []MessageElement{
				&MessageElementAt{
					Type: "test",
				},
			},
		},
		{
			Raw: `<sharp id="1" />`,
			Elements: []MessageElement{
				&MessageElementSharp{
					Id: "1",
				},
			},
		},
		{
			Raw: `<sharp id="1" name="test" />`,
			Elements: []MessageElement{
				&MessageElementSharp{
					Id:   "1",
					Name: "test",
				},
			},
		},
		{
			Raw: `<a href="https://example.com" />`,
			Elements: []MessageElement{
				&MessageElementA{
					Href: "https://example.com",
				},
			},
		},
	}
}
