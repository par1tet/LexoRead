import { Injectable, NestMiddleware } from '@nestjs/common';
import { Request, Response, NextFunction } from 'express';
import * as jwt from 'jsonwebtoken';
@Injectable()
export class UserMiddleware implements NestMiddleware {
  use(req: Request, res: Response, next: NextFunction) {
    const secret = process.env.JWT_SECRET_KEY;
    console.log(secret);

    if (req.method == 'OPTIONS') {
     return next();
    }
    const token = jwt.sign({ userId: 1 }, process.env.JWT_SECRET_KEY, { expiresIn: '5m' });
    if (!token) {
      res.status(403).send({ msg: 'пользователь не авторизован' });
    }
    console.log(token);
    
      const decodedData = jwt.verify(token, secret)   
      if(!decodedData) {
        res.status(403).send({msg: 'не валидный'})
      }
      req.body.user = decodedData
      next()
     
  }
}
