import { ApiHeader, ApiOperation, ApiProperty, ApiTags } from '@nestjs/swagger';
import { CHAR, INTEGER } from 'sequelize';
import { HasMany, Model } from 'sequelize-typescript';
import { Column, DataType, Table } from 'sequelize-typescript';
import { FavBooks } from './favBooks.model';
interface UserAttrs {
  isBanned: boolean;
  username: string;
  email: string;
  hashPassword: string;
  avatarFileUrl: string;
}
@Table({ tableName: 'users' })
export class User extends Model<User, UserAttrs> {
  @ApiProperty({ example: 1, description: 'Уникальный индификатор' })
  @Column({
    type: DataType.INTEGER,
    primaryKey: true,
    autoIncrement: true,
    unique: true,
  })
  id: number;
  @ApiProperty({ example: true, description: 'забанен' })
  @Column({ type: DataType.BOOLEAN, defaultValue: false, allowNull: true })
  isBanned: boolean;
  @ApiProperty({ example: 'Artem', description: 'username' })
  @Column({ type: DataType.STRING, allowNull: true })
  username: string;
  @ApiProperty({ example: 'smallgigached@gmail.com', description: 'email' })
  @Column({ type: DataType.STRING, unique: true })
  email: string;
  @ApiProperty({example: 'sdfsdfsdgfdf', description: 'JWT TOKEN User'})
  @Column({type: DataType.TEXT, allowNull: true})
  jwtToken: string
  @ApiProperty({
    example: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9',
    description: 'HashPassword',
  })
  @Column({ type: DataType.TEXT, allowNull: true })
  hashPassword: string;
  @ApiProperty({
    example:
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQy1QKMwvvFRtE7kNDp0oyv8e5k39uBdPxbbg&s',
    description: 'avatarUrl',
  })
  @Column({ type: DataType.STRING(500), allowNull: true })
  avatarFileUrl: string;
  @ApiProperty({ example: 'Hello World', description: 'description' })
  @Column({ type: DataType.STRING(350) })
  description: string;
  @HasMany(() => FavBooks)
  favBooks_Id: FavBooks[]
}
