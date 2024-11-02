import { Injectable } from '@nestjs/common'
import { InjectModel } from '@nestjs/sequelize'
import { Bookmark } from './bookmark.model'
import { CreateBookmarkDto } from './dto/create-bookmark.dto'

@Injectable()
export class BookmarkService {
	constructor(
		@InjectModel(Bookmark) private bookmarkRepository: typeof Bookmark
	) {}

	async createBookmark(dto: CreateBookmarkDto) {
		const bookmark = await this.bookmarkRepository.create(dto)
		return bookmark
	}

	async getAllBookmarks() {
		const bookmarks = await this.bookmarkRepository.findAll({
			include: { all: true },
		})
		return bookmarks
	}
}
