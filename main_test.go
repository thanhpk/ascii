package ascii

import (
	"fmt"
	"testing"
)

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert("越南社会主义共和国。独立与自由——幸福 Cộng hòa xã hội chủ nghĩa Việt Nam. Độc lập tự do - hạnh phúc 越南社会主义共和国。独立与自由——幸福 Cộng hòa xã hội chủ nghĩa Việt Nam. Độc lập tự do - hạnh phúc 越南社会主义共和国。独立与自由——幸福 Cộng hòa xã hội chủ nghĩa Việt Nam. Độc lập tự do - hạnh phúc 越南社会主义共和国。独立与自由——幸福 Cộng hòa xã hội chủ nghĩa Việt Nam. Độc lập tự do - hạnh phúc 越南社会主义共和国。独立与自由——幸福 Cộng hòa xã hội chủ nghĩa Việt Nam. Độc lập tự do - hạnh phúc 越南社会主义共和国。独立与自由——幸福 Cộng hòa xã hội chủ nghĩa Việt Nam. Độc lập tự do - hạnh phúc 越南社会主义共和国。独立与自由——幸福 Cộng hòa xã hội chủ nghĩa Việt Nam. Độc lập tự do - hạnh phúc 越南社会主义共和国。独立与自由——幸福 Cộng hòa xã hội chủ nghĩa Việt Nam. Độc lập tự do - hạnh phúc 越南社会主义共和国。独立与自由——幸福 Cộng hòa xã hội chủ nghĩa Việt Nam. Độc lập tự do - hạnh phúc 越南社会主义共和国。独立与自由——幸福 Cộng hòa xã hội chủ nghĩa Việt Nam. Độc lập tự do - hạnh phúc 越南社会主义共和国。独立与自由——幸福 Cộng hòa xã hội chủ nghĩa Việt Nam. Độc lập tự do - hạnh phúc 越南社会主义共和国。独立与自由——幸福 Cộng hòa xã hội chủ nghĩa Việt Nam. Độc lập tự do - hạnh phúc")
	}
}

func TestConvert(t *testing.T) {
	testcases := []struct {
		in     string
		expect string
	}{
		{"ậậậậ", "aaaa"},
		{"Bằng Minh Tuấn", "Bang Minh Tuan"},
		{"ß", ""},
		{"한글", ""},
		{"æ", ""},
		{"イーブイ", ""},
		{"Cộng hòa xã hội chủ nghĩa Việt Nam. Độc lập tự do - hạnh phúc", "Cong hoa xa hoi chu nghia Viet Nam. Doc lap tu do - hanh phuc"},
		{"République socialiste du Vietnam. Indépendance et liberté - bonheur", "Republique socialiste du Vietnam. Independance et liberte - bonheur"},
		{"Vietnam Sosyalist Cumhuriyeti. Bağımsızlık ve özgürlük - mutluluk", "Vietnam Sosyalist Cumhuriyeti. Bamszlk ve zgrlk - mutluluk"},
		{"Социјалистичке Републике Вијетнам. Независност и слобода - срећа", "  .    - "},
		{"越南社会主义共和国。独立与自由——幸福", ""},
	}

	for _, tc := range testcases {
		out := Convert(tc.in)
		if out != tc.expect {
			t.Errorf("expect \"%s\", got \"%s\"", tc.expect, out)
		}
	}
}

func ExampleCovnert() {
	fmt.Println(Convert("Nam quốc sơn hà mam đế cư"))
	fmt.Println(Convert("Cuántos años tienes"))
	fmt.Println(Convert("AppleInc 是苹果公司."))
	// Output:
	// Nam quoc son ha mam de cu
	// Cuantos aos tienes
	// AppleInc .
}
