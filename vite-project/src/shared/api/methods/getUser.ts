import axios from 'axios'
import {userService} from "../serviceLinks.ts";

export async function getUser(id: number) {
    console.log(id)
    let returnData = {}
    await axios.get(userService(`user/id/${id}`))
        .then(r => (returnData = r.data))
    returnData
}