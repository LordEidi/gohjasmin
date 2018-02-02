package model
/*-----------------------------------------------------------------------------
 **
 ** - GohJasmin -
 **
 ** Copyright 2017-18 by SwordLord - the coding crew - http://www.swordlord.com
 ** and contributing authors
 **
 ** This program is free software; you can redistribute it and/or modify it
 ** under the terms of the GNU Affero General Public License as published by the
 ** Free Software Foundation, either version 3 of the License, or (at your option)
 ** any later version.
 **
 ** This program is distributed in the hope that it will be useful, but WITHOUT
 ** ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 ** FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License
 ** for more details.
 **
 ** You should have received a copy of the GNU Affero General Public License
 ** along with this program. If not, see <http://www.gnu.org/licenses/>.
 **
 **-----------------------------------------------------------------------------
 **
 ** Original Authors:
 ** LordEidi@swordlord.com
 ** LordLightningBolt@swordlord.com
 **
-----------------------------------------------------------------------------*/

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Permission struct {
	PermissionId    int `gorm:"primary_key;AUTO_INCREMENT"`
	Permission 	string
	User string
	UserFK User `gorm:"ForeignKey:Name;AssociationForeignKey:UserId"`
	CrtDat	time.Time `sql:"DEFAULT:current_timestamp"`
	UpdDat	time.Time `sql:"DEFAULT:current_timestamp"`
}

func (m *Permission) BeforeUpdate(scope *gorm.Scope) (err error) {

	scope.SetColumn("UpdDat", time.Now())
	return  nil
}
