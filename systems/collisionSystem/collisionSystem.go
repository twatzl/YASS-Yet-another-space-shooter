package collisionSystem

type CollisionSystem interface {
	RegisterCircularCollider(c CircularCollider)
	UnregisterCircularCollider(c CircularCollider)
	RegisterTerrainCollider(c TerrainCollider)
}

type CollisionSystemControl interface {
	CollisionSystem
	Update()
}
