import { ApiProperty } from '@nestjs/swagger';
import { INTEGER } from 'sequelize';
import { Column, DataType, Model, Table } from 'sequelize-typescript';

interface UserAttrs {
  isBanned: boolean;
  username: string;
  email: string;
  hashPassword: string;
  avatarFileUrl: string;
  ListBook: number[];

}

@Table({ modelName: 'users'})
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
  @Column({ type: DataType.BOOLEAN, defaultValue: false})
  isBanned: boolean;
  @ApiProperty({ example: 'Artem', description: 'username' })
  @Column({ type: DataType.STRING, allowNull: true })
  username: string;
  @ApiProperty({ example: 'smallgigached@gmail.com', description: 'email' })
  @Column({ type: DataType.STRING, unique: true })
  email: string;
  @ApiProperty({
    example: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9',
    description: 'HashPassword',
  })
  @Column({ type: DataType.STRING, allowNull: true })
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
  @Column({ type: DataType.ARRAY(INTEGER) })
  ListBook: number[];
  
}
