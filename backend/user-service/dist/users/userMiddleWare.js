"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.UserMiddleware = void 0;
const common_1 = require("@nestjs/common");
const jwt = require("jsonwebtoken");
let UserMiddleware = class UserMiddleware {
    use(req, res, next) {
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
        const decodedData = jwt.verify(token, secret);
        if (!decodedData) {
            res.status(403).send({ msg: 'не валидный' });
        }
        req.body.user = decodedData;
        next();
    }
};
exports.UserMiddleware = UserMiddleware;
exports.UserMiddleware = UserMiddleware = __decorate([
    (0, common_1.Injectable)()
], UserMiddleware);
//# sourceMappingURL=userMiddleWare.js.map