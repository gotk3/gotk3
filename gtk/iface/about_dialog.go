package iface

import gdk_iface "github.com/gotk3/gotk3/gdk/iface"

type AboutDialog interface {
    Dialog

    AddCreditSection(string, []string)
    GetArtists() []string
    GetAuthors() []string
    GetComments() string
    GetCopyright() string
    GetDocumenters() []string
    GetLicense() string
    GetLicenseType() License
    GetLogoIconName() string
    GetProgramName() string
    GetTranslatorCredits() string
    GetVersion() string
    GetWebsite() string
    GetWebsiteLabel() string
    GetWrapLicense() bool
    SetArtists([]string)
    SetAuthors([]string)
    SetComments(string)
    SetCopyright(string)
    SetDocumenters([]string)
    SetLicense(string)
    SetLicenseType(License)
    SetLogo(gdk_iface.Pixbuf)
    SetLogoIconName(string)
    SetProgramName(string)
    SetTranslatorCredits(string)
    SetVersion(string)
    SetWebsite(string)
    SetWebsiteLabel(string)
    SetWrapLicense(bool)
} // end of AboutDialog

func AssertAboutDialog(_ AboutDialog) {}
