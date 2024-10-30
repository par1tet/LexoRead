import { Model } from 'sequelize-typescript';
import { FavBooks } from './favBooks.model';
interface UserAttrs {
    isBanned: boolean;
    username: string;
    email: string;
    hashPassword: string;
    avatarFileUrl: string;
}
export declare class User extends Model<User, UserAttrs> {
    id: number;
    isBanned: boolean;
    username: string;
    email: string;
    jwtToken: string;
    hashPassword: string;
    avatarFileUrl: string;
    description: string;
    favBooks_Id: FavBooks[];
}
export {};
