package messages

import (
	"../../npc"
	"../../tools"
	"../../types"
	"../../worldgen"
)

func WorldInfo (world *types.World) types.Message {
	var outMsg types.Message
	outMsg.Type = 0x07
	
	payload := tools.Int32ToBytes(int(world.Time))	
	var bb types.BitsByte = 0
	bb.SetBit(0, world.DayTime)
	bb.SetBit(1, world.BloodMoon)
	bb.SetBit(2, world.Eclipse)
	payload = append(payload, byte(bb))
	payload = append(payload, tools.Int8ToBytes(world.MoonPhase) ...)
	payload = append(payload, tools.Int16ToBytes(world.MaxTilesX) ...)
	payload = append(payload, tools.Int16ToBytes(world.MaxTilesY) ...)
	payload = append(payload, tools.Int16ToBytes(world.SpawnTileX) ...)
	payload = append(payload, tools.Int16ToBytes(world.SpawnTileY) ...)
	payload = append(payload, tools.Int16ToBytes(int(world.WorldSurface)) ...)
	payload = append(payload, tools.Int16ToBytes(int(world.RockLayer)) ...)
	payload = append(payload, tools.Int32ToBytes(world.WorldID) ...)
	payload = append(payload, tools.GetEncodedString(world.WorldName) ...)
	payload = append(payload, tools.Int8ToBytes(world.MoonType) ...)
	payload = append(payload, tools.Int8ToBytes(worldgen.TreeBG) ...)
	payload = append(payload, tools.Int8ToBytes(worldgen.CorruptBG) ...)
	payload = append(payload, tools.Int8ToBytes(worldgen.JungleBG) ...)
	payload = append(payload, tools.Int8ToBytes(worldgen.SnowBG) ...)
	payload = append(payload, tools.Int8ToBytes(worldgen.HallowBG) ...)
	payload = append(payload, tools.Int8ToBytes(worldgen.CrimsonBG) ...)
	payload = append(payload, tools.Int8ToBytes(worldgen.DesertBG) ...)
	payload = append(payload, tools.Int8ToBytes(worldgen.OceanBG) ...)
	payload = append(payload, tools.Int8ToBytes(world.IceBackStyle) ...)
	payload = append(payload, tools.Int8ToBytes(world.JungleBackStyle) ...)
	payload = append(payload, tools.Int8ToBytes(world.HellBackStyle) ...)
	payload = append(payload, tools.Int8ToBytes(world.MoonPhase) ...)
	payload = append(payload, tools.Float32ToBytes(world.WindSpeedSet) ...)
	payload = append(payload, tools.Int8ToBytes(world.NumClouds) ...)
	for i := 0; i < 3; i++ {
		payload = append(payload, tools.Int8ToBytes(world.TreeX[i]) ...)
	}
	for i := 0; i < 3; i++ {
		payload = append(payload, tools.Int8ToBytes(world.TreeStyle[i]) ...)
	}
	for i := 0; i < 3; i++ {
		payload = append(payload, tools.Int8ToBytes(world.CaveBackX[i]) ...)
	}
	for i := 0; i < 3; i++ {
		payload = append(payload, tools.Int8ToBytes(world.CaveBackStyle[i]) ...)
	}
	if !world.Raining {
		world.MaxRaining = 0.
	}
	payload = append(payload, tools.Float32ToBytes(world.MaxRaining) ...)
	bb = 0
	bb.SetBit(0, worldgen.ShadowOrbSmashed)
	bb.SetBit(1, npc.DownedBoss1)
	bb.SetBit(2, npc.DownedBoss2)
	bb.SetBit(3, npc.DownedBoss3)
	bb.SetBit(4, world.HardMode)
	bb.SetBit(5, npc.DownedClown)
	bb.SetBit(6, world.ServerSideCharacter)
	bb.SetBit(7, npc.DownedPlantBoss)
	payload = append(payload, byte(bb))
	bb = 0
	bb.SetBit(0, npc.DownedMechBoss1)
	bb.SetBit(1, npc.DownedMechBoss2)
	bb.SetBit(2, npc.DownedMechBoss3)
	bb.SetBit(3, npc.DownedMechBossAny)
	bb.SetBit(4, (world.CloudBGActive >= 1.))
	bb.SetBit(5, worldgen.Crimson)
	bb.SetBit(6, world.PumpkinMoon)
	bb.SetBit(7, world.SnowMoon)
	payload = append(payload, byte(bb))
	
	outMsg.Payload = payload
	outMsg.Length = int16(len(outMsg.Payload) + 3)
	return outMsg
}