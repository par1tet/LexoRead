import { Model } from 'sequelize-typescript';
import { User } from './user.model';
interface FavBooksAttrs {
    List: number;
}
export declare class FavBooks extends Model<FavBooks, FavBooksAttrs> {
    id: number;
    List: number;
    user: User;
    user_id: number;
}
export {};
