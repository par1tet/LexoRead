import { ApiProperty } from "@nestjs/swagger";

export class BanUserDto {
	@ApiProperty({example: 1})
  id: number;
}
