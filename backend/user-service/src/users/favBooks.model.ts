import { INTEGER } from 'sequelize';
import {
  BelongsTo,
  Column,
  DataType,
  ForeignKey,
  Model,
  Table,
} from 'sequelize-typescript';
import { User } from './user.model';
import { ApiProperty } from '@nestjs/swagger';
interface FavBooksAttrs {
  List: number;
}
@Table({ tableName: 'booksFavorite' })
export class FavBooks extends Model<FavBooks, FavBooksAttrs> {
  @ApiProperty({ example: 1, description: 'Уникальный индификатор' })
  @Column({ type: DataType.INTEGER, autoIncrement: true, primaryKey: true })
  id: number;
  @ApiProperty({
    example: [1, 2, 3, 4, 5],
    description: 'получение книг по Id',
  })
  @Column({ type: DataType.ARRAY(INTEGER), allowNull: true })
  List: number;
  @BelongsTo(() => User)
  user: User;
  @ForeignKey(() => User)
  user_id: number;
}
