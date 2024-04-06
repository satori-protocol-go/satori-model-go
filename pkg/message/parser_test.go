package message

import (
	"testing"
)

/**
group('Elements', () {
    test('Basic', () {
      expect(n.parse('<at id="1"/>'), equals([n.At(id: '1')]));
      expect(n.parse('<at id="1" name="test"/>'),
          equals([n.At(id: '1', name: 'test')]));
      expect(n.parse('<at role="test"/>'), equals([n.At(role: 'test')]));
      expect(n.parse('<at type="test"/>'), equals([n.At(type: 'test')]));
      expect(n.parse('<sharp id="1"/>'), equals([n.Sharp(id: '1')]));
      expect(n.parse('<sharp id="1" name="test"/>'),
          equals([n.Sharp(id: '1', name: 'test')]));
      expect(n.parse('<a href="https://example.com"/>'),
          equals([n.Anchor(href: 'https://example.com')]));
    });

    test('Layout', () {
      expect(n.parse('<br/>'), equals([n.LineBreak()]));
      expect(n.parse('<p/>'), equals([n.Paragraph()]));
      expect(
          n.parse('<p>test</p>'),
          equals([
            n.Paragraph(children: [n.Text('test')])
          ]));
      expect(
          n.parse('<p>testtest2</p>'),
          equals([
            n.Paragraph(children: [n.Text('testtest2')])
          ]));
      expect(
          n.parse('<p>test<br/></p>'),
          equals([
            n.Paragraph(children: [n.Text('test'), n.LineBreak()])
          ]));
      expect(n.parse('<message/>'), equals([n.Message()]));
      expect(n.parse('<message id="1"/>'), equals([n.Message(id: '1')]));
      expect(n.parse('<message forward/>'), equals([n.Message(forward: true)]));
      expect(n.parse('<message id="1" forward/>'),
          equals([n.Message(id: '1', forward: true)]));
      // NOTE: Should we handle `key="false"`?
      // expect(n.parse('<message id="1" forward="false"/>'),
      //     equals([n.Message(id: '1', forward: false)]));
      expect(
          n.parse('<message>test</message>'),
          equals([
            n.Message(children: [n.Text('test')])
          ]));
    });

    test('Meta', () {
      expect(n.parse('<quote/>'), equals([n.Quote()]));
      expect(
          n.parse('<quote>test</quote>'),
          equals([
            n.Quote(children: [n.Text('test')])
          ]));
      expect(n.parse('<author/>'), equals([n.Author()]));
      expect(n.parse('<author user-id="1"/>'), equals([n.Author(userId: '1')]));
      expect(n.parse('<author user-id="1" nickname="test"/>'),
          equals([n.Author(userId: '1', nickname: 'test')]));
      expect(n.parse('<author user-id="1" nickname="test" avatar="test"/>'),
          equals([n.Author(userId: '1', nickname: 'test', avatar: 'test')]));
    });

    test('Resource', () {
      expect(n.parse('<img src="https://example.com"/>'),
          equals([n.Image(src: 'https://example.com')]));
      expect(
          n.parse('<img src="https://example.com" cache timeout="1000"/>'),
          equals([
            n.Image(src: 'https://example.com', cache: true, timeout: '1000')
          ]));
      expect(
          n.parse('<img src="https://example.com" width="300" height="200"/>'),
          equals(
              [n.Image(src: 'https://example.com', width: 300, height: 200)]));
      expect(n.parse('<audio src="https://example.com"/>'),
          equals([n.Audio(src: 'https://example.com')]));
      expect(n.parse('<video src="https://example.com"/>'),
          equals([n.Video(src: 'https://example.com')]));
    });

    test('Modifier', () {
      expect(n.parse('<strong>test</strong>'), equals([n.Bold('test')]));
      expect(n.parse('<b>test</b>'), equals([n.Bold('test')]));
      expect(n.parse('<em>test</em>'), equals([n.Idiomatic('test')]));
      expect(n.parse('<i>test</i>'), equals([n.Idiomatic('test')]));
      expect(n.parse('<ins>test</ins>'), equals([n.Underline('test')]));
      expect(n.parse('<u>test</u>'), equals([n.Underline('test')]));
      expect(n.parse('<del>test</del>'), equals([n.Strikethrough('test')]));
      expect(n.parse('<s>test</s>'), equals([n.Strikethrough('test')]));
      expect(n.parse('<spl>test</spl>'), equals([n.Spoiler('test')]));
      expect(n.parse('<code>test</code>'), equals([n.Code('test')]));
      expect(n.parse('<sup>test</sup>'), equals([n.Superscript('test')]));
      expect(n.parse('<sub>test</sub>'), equals([n.Subscript('test')]));
    });
  });

  group('Examples', () {
    test('Single Message', () {
      expect(
          n.parse("""<message>
  <at id="1" name="test"/>
  <at id="2" name="test2"/>
  <br/>
  <p>test</p>
  <img src="https://example.com"/>
  <audio src="https://example.com"/>
  <p>
    <b>bold</b><sup>sup</sup><sub>sub</sub>
    <spl>spoiler</spl>
    <code>code</code>
    <u>underline</u>
    <s>strikethrough</s>
  </p>
</message>"""),
          equals([
            n.Message(
              children: [
                n.At(id: '1', name: 'test'),
                n.At(id: '2', name: 'test2'),
                n.LineBreak(),
                n.Paragraph(children: [n.Text('test')]),
                n.Image(src: 'https://example.com'),
                n.Audio(src: 'https://example.com'),
                n.Paragraph(children: [
                  n.Bold('bold'),
                  n.Superscript('sup'),
                  n.Subscript('sub'),
                  n.Spoiler('spoiler'),
                  n.Code('code'),
                  n.Underline('underline'),
                  n.Strikethrough('strikethrough')
                ])
              ],
            )
          ]));

      expect(
          n.parse("""<message>
<code>
  yarn add nitori
  yarn run test
  npm publish
</code>
</message>"""),
          equals([
            n.Message(
              children: [
                n.Code('yarn add nitori\n  yarn run test\n  npm publish')
              ],
            )
          ]));
    });
  });
*/
//raw_message["baisc"] = []string"纯文本"}

type rawToElement struct {
	Raw      string
	Elements []MessageElement
}

func getRawMessage() map[string][]rawToElement {
	raw_message := make(map[string][]rawToElement)
	raw_message["basic"] = []rawToElement{
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
	raw_message["Layout"] = []rawToElement{
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
					childrenMessageElement: &childrenMessageElement{
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
					childrenMessageElement: &childrenMessageElement{
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
					childrenMessageElement: &childrenMessageElement{
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
	}
	return raw_message
}

func Test(t *testing.T) {
	for group, messages := range getRawMessage() {
		t.Logf("start test group: %s", group)
		for _, message := range messages {
			_, err := Parse(message.Raw)
			if err != nil {
				t.Fatalf("%s Parse error: %s", message.Raw, err)
			}
			result, err := Stringify(message.Elements)
			if err != nil {
				t.Fatalf("%s Stringify error: %s", message.Elements, err)
			}
			if result != message.Raw {
				t.Fatalf("%s not eq %s", result, message.Raw)
			}
		}
	}
}
