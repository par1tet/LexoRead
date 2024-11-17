import {User} from '../../src/entities/classes/user.ts'

export function createFriendList(coverUrls: string[]): User[] {
    const friendsArrays: User[] = []

    coverUrls.forEach((coverUrl) => {
        friendsArrays.push(new User(coverUrl));
    })
    return friendsArrays
}
