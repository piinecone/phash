#include <pHash.h>

#ifdef __cplusplus
extern "C" {
#endif

void
cimg_exception_mode_quiet() {
	cimg::exception_mode(0);
}

int
ph_dct_imagehash_wrapper(const char *file, ulong64 *hash) {
	int result = ph_dct_imagehash(file, *hash);
	if (result == 0) {
		errno = 0;
	}
	return result;
}

#ifdef __cplusplus
}
#endif

