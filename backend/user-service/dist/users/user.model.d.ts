import { Model } from 'sequelize-typescript';
interface UserAttrs {
    isBanned: boolean;
    username: string;
    email: string;
    hashPassword: string;
    avatarFileUrl: string;
    ListBook: number[];
}
export declare class User extends Model<User, UserAttrs> {
    id: number;
    isBanned: boolean;
    username: string;
    email: string;
    hashPassword: string;
    avatarFileUrl: string;
    description: string;
    ListBook: number[];
}
export {};
