/*
 * Project Polar
 *
 * http://iusworks.com/project/polar
 * hongl.gu@iusworks.com
 *
 *  Copyright (c)  2021. The Iusworks Network, LLC.  All rights reserved.
 *
 */

package message

import "time"

type Health struct {
	Health string    `json:"health" xml:"health" yaml:"health"`
	Time   time.Time `json:"time" xml:"time" yaml:"time"`
}

func Healthy() *Health {
	return &Health{
		Health: "fine",
		Time:   time.Now(),
	}
}
