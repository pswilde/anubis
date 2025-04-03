// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package web

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/TecharoHQ/anubis"
	"github.com/TecharoHQ/anubis/xess"
)

func base(title string, body templ.Component) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html><head><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 12, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</title><link rel=\"stylesheet\" href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(xess.URL)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 13, Col: 41}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><style>\n    body,\n    html {\n      height: 100%;\n      display: flex;\n      justify-content: center;\n      align-items: center;\n      margin-left: auto;\n      margin-right: auto;\n    }\n\n    .centered-div {\n      text-align: center;\n    }\n\n    .lds-roller,\n    .lds-roller div,\n    .lds-roller div:after {\n      box-sizing: border-box;\n    }\n\n    .lds-roller {\n      display: inline-block;\n      position: relative;\n      width: 80px;\n      height: 80px;\n    }\n\n    .lds-roller div {\n      animation: lds-roller 1.2s cubic-bezier(0.5, 0, 0.5, 1) infinite;\n      transform-origin: 40px 40px;\n    }\n\n    .lds-roller div:after {\n      content: \" \";\n      display: block;\n      position: absolute;\n      width: 7.2px;\n      height: 7.2px;\n      border-radius: 50%;\n      background: currentColor;\n      margin: -3.6px 0 0 -3.6px;\n    }\n\n    .lds-roller div:nth-child(1) {\n      animation-delay: -0.036s;\n    }\n\n    .lds-roller div:nth-child(1):after {\n      top: 62.62742px;\n      left: 62.62742px;\n    }\n\n    .lds-roller div:nth-child(2) {\n      animation-delay: -0.072s;\n    }\n\n    .lds-roller div:nth-child(2):after {\n      top: 67.71281px;\n      left: 56px;\n    }\n\n    .lds-roller div:nth-child(3) {\n      animation-delay: -0.108s;\n    }\n\n    .lds-roller div:nth-child(3):after {\n      top: 70.90963px;\n      left: 48.28221px;\n    }\n\n    .lds-roller div:nth-child(4) {\n      animation-delay: -0.144s;\n    }\n\n    .lds-roller div:nth-child(4):after {\n      top: 72px;\n      left: 40px;\n    }\n\n    .lds-roller div:nth-child(5) {\n      animation-delay: -0.18s;\n    }\n\n    .lds-roller div:nth-child(5):after {\n      top: 70.90963px;\n      left: 31.71779px;\n    }\n\n    .lds-roller div:nth-child(6) {\n      animation-delay: -0.216s;\n    }\n\n    .lds-roller div:nth-child(6):after {\n      top: 67.71281px;\n      left: 24px;\n    }\n\n    .lds-roller div:nth-child(7) {\n      animation-delay: -0.252s;\n    }\n\n    .lds-roller div:nth-child(7):after {\n      top: 62.62742px;\n      left: 17.37258px;\n    }\n\n    .lds-roller div:nth-child(8) {\n      animation-delay: -0.288s;\n    }\n\n    .lds-roller div:nth-child(8):after {\n      top: 56px;\n      left: 12.28719px;\n    }\n\n    @keyframes lds-roller {\n      0% {\n        transform: rotate(0deg);\n      }\n\n      100% {\n        transform: rotate(360deg);\n      }\n    }\n      </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.JSONScript("anubis_version", anubis.Version).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</head><body id=\"top\"><main><center><h1 id=\"title\" class=\".centered-div\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 146, Col: 49}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</h1></center>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = body.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<footer><center><p>Protected by <a href=\"https://github.com/TecharoHQ/anubis\">Anubis</a> from <a href=\"https://techaro.lol\">Techaro</a>. Made with ❤️ in 🇨🇦.</p><p>Mascot design by <a href=\"https://bsky.app/profile/celphase.bsky.social\">CELPHASE</a>.</p></center></footer></main></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func index() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<div class=\"centered-div\"><img id=\"image\" style=\"width:100%;max-width:256px;\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs("/.within.website/x/cmd/anubis/static/img/pensive.webp?cacheBuster=" +
			anubis.Version)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 170, Col: 18}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\"> <img style=\"display:none;\" style=\"width:100%;max-width:256px;\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs("/.within.website/x/cmd/anubis/static/img/happy.webp?cacheBuster=" +
			anubis.Version)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 176, Col: 18}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\"><p id=\"status\">Loading...</p><script async type=\"module\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs("/.within.website/x/cmd/anubis/static/js/main.mjs?cacheBuster=" + anubis.Version)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 179, Col: 116}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\"></script><div id=\"spinner\" class=\"lds-roller\"><div></div><div></div><div></div><div></div><div></div><div></div><div></div><div></div></div><details><summary>Why am I seeing this?</summary><p>You are seeing this because the administrator of this website has set up <a href=\"https://github.com/TecharoHQ/anubis\">Anubis</a> to protect the server against the scourge of <a href=\"https://thelibre.news/foss-infrastructure-is-under-attack-by-ai-companies/\">AI companies aggressively scraping websites</a>. This can and does cause downtime for the websites, which makes their resources inaccessible for everyone.</p><p>Anubis is a compromise. Anubis uses a <a href=\"https://anubis.techaro.lol/docs/design/why-proof-of-work\">Proof-of-Work</a> scheme in the vein of <a href=\"https://en.wikipedia.org/wiki/Hashcash\">Hashcash</a>, a proposed proof-of-work scheme for reducing email spam. The idea is that at individual scales the additional load is ignorable, but at mass scraper levels it adds up and makes scraping much more expensive.</p><p>Ultimately, this is a hack whose real purpose is to give a \"good enough\" placeholder solution so that more time can be spent on fingerprinting and identifying headless browsers (EG: via how they do font rendering) so that the challenge proof of work page doesn't need to be presented to users that are much more likely to be legitimate.</p><p>Please note that Anubis requires the use of modern JavaScript features that plugins like <a href=\"https://jshelter.org/\">JShelter</a> will disable. Please disable JShelter or other such plugins for this domain.</p></details><noscript><p>Sadly, you must enable JavaScript to get past this challenge. This is required because AI companies have changed the social contract around how website hosting works. A no-JS solution is a work-in-progress.</p></noscript><div id=\"testarea\"></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func errorPage(message string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<div class=\"centered-div\"><img id=\"image\" style=\"width:100%;max-width:256px;\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs("/.within.website/x/cmd/anubis/static/img/reject.webp?cacheBuster=" + anubis.Version)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 212, Col: 93}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "\"><p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(message)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `index.templ`, Line: 214, Col: 14}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, ".</p><button onClick=\"window.location.reload();\">Try again</button><p><a href=\"/\">Go home</a></p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
