/**
 * Interacting with userinfo in localStorage
 * @constructor
 * @return {UserInfo}
 * json(): json of UserInfo
 * remove(): remove UserInfo from localStorage
 * set(): set string of UserInfo to localStorage 
 */
function UserInfo() {
    const USER_INFO_ID = "userinfo"

    return {
        json: () => { return JSON.parse(localStorage.getItem(USER_INFO_ID)) },
        remove: () => localStorage.removeItem(USER_INFO_ID),
        set:(info="")=>localStorage.setItem("userinfo", info),
    }
}