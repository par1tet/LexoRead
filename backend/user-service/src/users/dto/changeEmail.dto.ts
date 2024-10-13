import { ApiProperty } from "@nestjs/swagger";

export class ChangeEmailDto {
	@ApiProperty({example: 1})
	userId: number;
	@ApiProperty({example: 'smallgigach@gmail.com'})
	newEmail: string;
}