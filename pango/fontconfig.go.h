
#include <fontconfig/fontconfig.h>

static int addFont(char* font) {
    FcBool fontAddStatus = FcConfigAppFontAddFile(FcConfigGetCurrent(), font);
    return fontAddStatus;
}