import { Module } from '@nestjs/common'
import { BookmarkController } from './bookmark.controller'
import { BookmarkService } from './bookmark.service'
import { SequelizeModule } from '@nestjs/sequelize'
import { Bookmark } from './bookmark.model'

@Module({
	controllers: [BookmarkController],
	providers: [BookmarkService],
	imports: [SequelizeModule.forFeature([Bookmark])],
	exports: [BookmarkService],
})
export class BookmarkModule {}
