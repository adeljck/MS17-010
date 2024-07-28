#include <windows.h>


DWORD EnableRDPServiceInternal(void)

{
    HKEY hKey = NULL;
    if (RegOpenKeyEx(
        HKEY_LOCAL_MACHINE,
        "SYSTEM\\CurrentControlSet\\Control\\Terminal Server",
        0,
        KEY_WRITE,
        &hKey) != ERROR_SUCCESS)
    {
        return -1;
    }

    // Set the value to enable RDP
    DWORD dwValue = 0;

    if (RegSetValueEx(
        hKey,
        "fDenyTSConnections",
        0,
        REG_DWORD,
        (LPBYTE)&dwValue, 
        sizeof(dwValue)) != ERROR_SUCCESS)
    {
        return -1;
    }
    else {
        RegCloseKey(hKey);
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
