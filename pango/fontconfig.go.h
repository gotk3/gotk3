
#include <fontconfig/fontconfig.h>

static int addFont(unsigned char *font) {
  FcBool fontAddStatus = FcConfigAppFontAddFile(FcConfigGetCurrent(), font);
  return fontAddStatus;
}