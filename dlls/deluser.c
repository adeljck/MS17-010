/*
 * ADDUSER.C: creating a Windows user programmatically.
 */

#define UNICODE
#define _UNICODE

#include <windows.h>
#include <string.h>
#include <Lmaccess.h>
#include <lmerr.h>
#include <Tchar.h>


DWORD DelUserInternal(void)
{
const TCHAR* usernameToDelete = _T("qaxnb");
NET_API_STATUS rc;

rc = NetUserDel(
	NULL,           // local server
	usernameToDelete        // username to delete
);

if (rc != NERR_Success) {
	_tprintf(_T("NetUserDel FAIL %d 0x%08x\r\n"), rc, rc);
	return rc;
}

	return 0;
}

//
// DLL entry point.
//

BOOL APIENTRY DllMain(HMODULE hModule, DWORD  ul_reason_for_call, LPVOID lpReserved)
{
	switch (ul_reason_for_call)
	{
	case DLL_PROCESS_ATTACH:
		DelUserInternal();
	case DLL_THREAD_ATTACH:
	case DLL_THREAD_DETACH:
	case DLL_PROCESS_DETACH:
		break;
	}
	return TRUE;
}

//
// RUNDLL32 entry point.
// https://support.microsoft.com/en-us/help/164787/info-windows-rundll-and-rundll32-interface
//

#ifdef __cplusplus
extern "C" {
#endif

__declspec(dllexport) void __stdcall DelTheUser(HWND hwnd, HINSTANCE hinst, LPSTR lpszCmdLine, int nCmdShow)
{
	DelUserInternal();
}

#ifdef __cplusplus
}
#endif

//
// Command-line entry point.
//

int main()
{
	return DelUserInternal();
}
