import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/sequelize';
import { Bookmark } from './bookmark.model';

@Injectable()
export class BookmarkService {
  constructor(@InjectModel(Bookmark) private bookmarkRepository: typeof Bookmark) {
    
  }
}
