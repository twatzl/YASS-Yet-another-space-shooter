package collisionSystem

type colliderSystemImpl struct {
	circularColliders map[CircularCollider]CircularCollider
	terrainCollider   TerrainCollider
}

func New() CollisionSystemControl {
	return &colliderSystemImpl{
		circularColliders: make(map[CircularCollider]CircularCollider),
		terrainCollider:   nil,
	}
}

func (cs *colliderSystemImpl) RegisterCircularCollider(c CircularCollider) {
	cs.circularColliders[c] = c
}

func (cs *colliderSystemImpl) UnregisterCircularCollider(c CircularCollider) {
	delete(cs.circularColliders, c)
}

func (cs *colliderSystemImpl) RegisterTerrainCollider(c TerrainCollider) {
	cs.terrainCollider = c
}

func (cs *colliderSystemImpl) Update() {
	/* check collisions of objects with terrain */
	tc := cs.terrainCollider
	for _, cc := range cs.circularColliders {
		if checkCollisionCircTerrain(tc, cc) {
			cc.callCollisionCallback(nil)
		}
	}

	/* check collisions of objects with each other */
	circularColliders := make([]CircularCollider, 0, len(cs.circularColliders))
	for k, _ := range cs.circularColliders {
		circularColliders = append(circularColliders, k)
	}

	for i := 0; i < len(circularColliders); i++ {
		for j := i + 1; j < len(circularColliders); j++ {
			cc1 := circularColliders[i]
			cc2 := circularColliders[j]
			if checkCollisionCircCirc(cc1, cc2) {
				cc1.callCollisionCallback(cc2.getGameObject())
				cc2.callCollisionCallback(cc1.getGameObject())
			}
		}
	}
}

func checkCollisionCircCirc(c1, c2 CircularCollider) bool {
	center1 := c1.GetCenter()
	center2 := c2.GetCenter()
	r1 := c1.GetRadius()
	r2 := c2.GetRadius()

	dist := center1.Sub(center2).Len()
	return dist < r1 + r2
}

func checkCollisionCircTerrain(tc TerrainCollider, cc CircularCollider) bool {
	// sin(45) = cos(45) is approx 0.7
	diagDist := int(0.7 * cc.GetRadius())
	radius := int(cc.GetRadius())
	center := cc.GetCenter()
	cx := int(center.X)
	cy := int(center.Y)

	for i := -radius; i <= radius; i++ {
		if tc.CollidesAt(i + cx, cy) || tc.CollidesAt(cx, i + cy) {
			return true
		}
	}

	for j := -diagDist; j <= diagDist; j++ {
		if tc.CollidesAt(j+cx, j+cy) || tc.CollidesAt(cx-j, j+cy) {
			return true
		}
	}
	return false
}
