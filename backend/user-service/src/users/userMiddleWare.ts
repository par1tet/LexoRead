import { Injectable, NestMiddleware } from '@nestjs/common';
import { Request, Response, NextFunction } from 'express';
import jwt from 'jsonwebtoken';

@Injectable()
export class UserMiddleware implements NestMiddleware {
  use(req: Request, res: Response, next: NextFunction) {
    const secret = process.env.JWT_SECRET_KEY;
    console.log('JWT Secret:', secret);

    if (req.method === 'OPTIONS') {
      return next();
    }

    const token = jwt.sign({ id: 1 }, secret);
    console.log('Generated Token:', token);

    const decoded = jwt.verify(token, secret);
    console.log('Decoded Token:', decoded);

    if (!token) {
      return res.status(403).send({ msg: 'пользователь не авторизован' });
    }

    try {
      const decodedData = jwt.verify(token, secret);
      console.log('Decoded Data:', decodedData);
      next();
    } catch (error) {
      console.error('JWT Verification Error:', error);
      return res.status(403).send({ msg: 'Invalid token' });
    }
  }
}
