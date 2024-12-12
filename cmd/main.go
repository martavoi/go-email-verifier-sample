package main

import (
	"fmt"
	"net/mail"
	"strings"

	emailverifier "github.com/AfterShip/email-verifier"
)

func main() {

	v := emailverifier.
		NewVerifier().
		EnableDomainSuggest()

	emails := [...]string{
		"quentonholloway65💯@gmail.com",
		"maria.mendozamendo…@gmail.com",
		"shabnamrizad۶۰@gmail.com",
		"jesúsarcoz@g.com",
		"🥱3rtwe@yahoo.comws.ijjjwkss",
		"pablomalaño@75.com",
		"cheerykeelouranià@gmail.com.thisisavalidemail",
		"¡35043422@gmail.com",
		"camposĺùis.2122-@mail.com",
		"wwwkervensjean\"12@gmail.com",
		"ᵢₙfₒ@nobelcareer.com",
		"b···········@gmail.com",
		"³720joshua@gmail.com",
		"a.h.j.ramón.2020@hotmail.con",
		"thaic6⁷57@gma8l.com",
		"gonzález843@iclud.con",
		"máquinaazulcruzazul3@g.mail",
		"román.alfredo@aiclud.con",
		"alexnuñez5101989@g.con",
		"yuri12.2030@gmail.acom",
		"jeanjacquesamyot7@gmail.comj",
		"merlebthornton20@gmail.comm",
		"sheilaramirez5444@gmail.c.om",
		"sunday3366@g.mail.com",
		"luyapang17591@hxm1.aleleikxkj.shop",
		"𝓂𝓅𝓇𝒶𝓅ℯ𝒹𝓇ℴ.𝓃ℴ𝓇𝒶𝓁ℯ𝓈12@𝓰𝓂𝒶𝒾𝓁.𝒸ℴ𝓂",
		"dzmitry.martavoi@gmail.com",
		"pm@martavoi.by",
	}

	for _, email := range emails {
		sb := strings.Builder{}

		_, err := mail.ParseAddress(email)
		if err == nil {
			sb.WriteString(fmt.Sprintf("RFC 5322 valid email string: %s ", email))
		} else {
			sb.WriteString(fmt.Sprintf("RFC 5322 invalid email string: %s ", email))
		}

		ret, err := v.Verify(email)

		if err != nil {
			sb.WriteString(fmt.Sprintf("Error: { %v }; ", err))
		}

		if ret.Disposable {
			sb.WriteString("Is Disposable Email; ")
		}

		if !ret.HasMxRecords {
			sb.WriteString("Doesn't have MX records; ")
		}

		if !ret.Syntax.Valid {
			sb.WriteString("Syntax is Invalid; ")
		}

		if len(ret.Suggestion) > 0 {
			sb.WriteString(fmt.Sprintf("Suggested domain is %s; ", ret.Suggestion))
		}

		fmt.Println(sb.String())
	}
}
