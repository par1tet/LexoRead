import { Column, DataType, Model } from "sequelize-typescript";



interface BookmarkCreationAttrs {
  content: string;
  numberPage: number;
}

export class Bookmark extends Model<Bookmark> {

  @Column({type: DataType.INTEGER, unique: true, autoIncrement: true, primaryKey: true})
  id: number;

  @Column({type: DataType.STRING})
  content: string;

  @Column({type: DataType.INTEGER })
  numberPage: number;

}