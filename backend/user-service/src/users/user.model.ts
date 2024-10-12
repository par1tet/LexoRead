import { CHAR } from 'sequelize';
import { Model } from 'sequelize-typescript';
import { Column, DataType, Table } from 'sequelize-typescript';
interface UserAttrs {
  isBanned: boolean;
  username: string;
  email: string;
  hashPassword: string;
  avatarFileUrl: string;
}
@Table({ tableName: 'users' })
export class User extends Model<User, UserAttrs> {
  @Column({
    type: DataType.INTEGER,
    primaryKey: true,
    autoIncrement: true,
    unique: true,
  })
  id: number;
  @Column({ type: DataType.BOOLEAN, defaultValue: false, allowNull: true})
  isBanned: boolean;
  @Column({ type: DataType.STRING, allowNull: true })
  username: string;
  @Column({ type: DataType.STRING, unique: true,  })
  email: string;
  @Column({ type: DataType.STRING, allowNull: true })
  hashPassword: string;
  @Column({ type: DataType.STRING(500), allowNull: true })
  avatarFileUrl: string;
  @Column({type: DataType.STRING(350)})
  description: string
}
