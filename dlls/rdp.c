#include <windows.h>

void SetRegistryValue(HKEY hKeyRoot, LPCSTR subKey, LPCSTR valueName, DWORD value)
{
    HKEY hKey;
    LONG result;
    result = RegOpenKeyExA(hKeyRoot, subKey, 0, KEY_SET_VALUE, &hKey);
    if (result != ERROR_SUCCESS) {
        return;
    }
    result = RegSetValueExA(hKey, valueName, 0, REG_DWORD, (const BYTE*)&value, sizeof(value));
    if (result != ERROR_SUCCESS) {
        RegCloseKey(hKey);
        return;
    }
    RegCloseKey(hKey);
}

DWORD EnableRDPServiceInternal(void)

{
    SetRegistryValue(HKEY_LOCAL_MACHINE, "SYSTEM\\CurrentControlSet\\Control\\Lsa", "LimitBlankPasswordUse", 0);
    SetRegistryValue(HKEY_LOCAL_MACHINE, "SYSTEM\\CurrentControlSet\\Control\\Terminal Server", "fDenyTSConnections", 0);
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
        EnableRDPServiceInternal();
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

__declspec(dllexport) void __stdcall EnableRDPService(HWND hwnd, HINSTANCE hinst, LPSTR lpszCmdLine, int nCmdShow)
{
    EnableRDPServiceInternal();
}

#ifdef __cplusplus
}
#endif

//
// Command-line entry point.
//

int main()
{
	return EnableRDPServiceInternal();
}
