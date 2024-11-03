import { IUser } from "../interfaces/IUser";

export class User implements IUser {
	constructor(avatarUrl: string, username: string) {
		this.avatarUrl = avatarUrl;
		this.username = username;
	}
	avatarUrl: string;
	username: string;
}