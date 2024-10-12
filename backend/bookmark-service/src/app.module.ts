import { Module } from "@nestjs/common";
import { BookmarkModule } from './bookmark/bookmark.module';
import { SequelizeModule } from "@nestjs/sequelize";
import { ConfigModule } from "@nestjs/config";


@Module({
	controllers: [],
	imports: [
		BookmarkModule,
		ConfigModule.forRoot({
			envFilePath: `.${process.env.NODE_ENV}.env`,
		}),
		SequelizeModule.forRoot({
			dialect: 'postgres',
			host: process.env.POSTGRES_HOST,
			port: Number(process.env.POSTGRES_PORT),
			username: process.env.POSTGRES_USER,
			password: process.env.POSTGRES_PASSWORD,
			database: process.env.POSTGRES_DB,
			autoLoadModels: true,
			synchronize: true,
			models: [],
		}),
	],
})
export class AppModule {}