import axios from "axios";
import {userService} from "../serviceLinks.ts";
import {User} from "../../../entities/classes/user.ts";
export async function getUser(userId: number | null): Promise<User> {
    console.log(userId)
    let userData = []

    await axios.get(userService(`user/id/${userId}`)).then((response) => {
        userData = response.data;
    })
    return userData;
}