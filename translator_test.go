package tts

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewTranslator(t *testing.T) {
	Convey("Attempting to consume bad xml should produce an error", t, func() {
		translator, err := NewTranslator("", "en")
		So(translator, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("A translator using acceptable data should produce expected results", t, func() {
		xmlData := `<translations>
	<en> <!-- English -->
		<strings>
            <push_message>testing</push_message>
            <sound_bae>Bae</sound_bae>
            <sound_baseball>Baseball</sound_baseball>
        </strings>
    </en>
</translations>`

		translator, err := NewTranslator(xmlData, "en")
		So(err, ShouldBeNil)

		So(translator.xmlData, ShouldEqual, xmlData)
		So(translator.defaultLanguageCode, ShouldEqual, "en")

		So(translator.Get("en", "push_message"), ShouldEqual, "testing")
		So(translator.Get("en", "sound_bae"), ShouldEqual, "Bae")
		So(translator.Get("en", "sound_baseball"), ShouldEqual, "Baseball")

		// An unknown language code should default to english
		So(translator.Get("es", "push_message"), ShouldEqual, "testing")
		So(translator.Get("es", "sound_bae"), ShouldEqual, "Bae")
		So(translator.Get("es", "sound_baseball"), ShouldEqual, "Baseball")
	})
}
