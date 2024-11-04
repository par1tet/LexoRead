import { Injectable, NestMiddleware } from '@nestjs/common';
import { Request, Response, NextFunction } from 'express';
import * as jwt from 'jsonwebtoken';
import 'core-js/stable/atob';
declare global {
  namespace Express {
    interface Request {
      user?: any;
    }
  }
}
@Injectable()
export class UserMiddleware implements NestMiddleware {
  use(req: Request, res: Response, next: NextFunction) {
    try {
      const secret = process.env.JWT_SECRET_KEY;
      if (req.method === 'OPTIONS') {
        return next();
      }
      const authHeader = req.headers.authorization;
      if (!authHeader) {
        return res.status(403).send({ msg: 'пользователь не авторизован' });
      }
      const token = authHeader.split(' ')[1];
      const decodedData = jwt.verify(token, secret, { algorithms: ['HS256'] });
      next();
      req.user = decodedData;
    } catch (err) {
      console.log(err);
    }
  }
}
