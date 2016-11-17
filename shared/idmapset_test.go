package shared

import (
	"fmt"
	"testing"
)

func TestIdmapSetAddSafe_split(t *testing.T) {
	orig := IdmapSet{Idmap: []IdmapEntry{IdmapEntry{Isuid: true, Hostid: 1000, Nsid: 0, Maprange: 1000}}}

	if err := orig.AddSafe(IdmapEntry{Isuid: true, Hostid: 500, Nsid: 500, Maprange: 10}); err != nil {
		t.Error(err)
		return
	}

	if orig.Idmap[0].Hostid != 1000 || orig.Idmap[0].Nsid != 0 || orig.Idmap[0].Maprange != 500 {
		t.Error(fmt.Errorf("bad range: %v", orig.Idmap[0]))
		return
	}

	if orig.Idmap[1].Hostid != 500 || orig.Idmap[1].Nsid != 500 || orig.Idmap[1].Maprange != 10 {
		t.Error(fmt.Errorf("bad range: %v", orig.Idmap[1]))
		return
	}

	if orig.Idmap[2].Hostid != 1510 || orig.Idmap[2].Nsid != 510 || orig.Idmap[2].Maprange != 490 {
		t.Error(fmt.Errorf("bad range: %v", orig.Idmap[2]))
		return
	}

	if len(orig.Idmap) != 3 {
		t.Error("too many idmap entries")
		return
	}
}

func TestIdmapSetAddSafe_lower(t *testing.T) {
	orig := IdmapSet{Idmap: []IdmapEntry{IdmapEntry{Isuid: true, Hostid: 1000, Nsid: 0, Maprange: 1000}}}

	if err := orig.AddSafe(IdmapEntry{Isuid: true, Hostid: 500, Nsid: 0, Maprange: 10}); err != nil {
		t.Error(err)
		return
	}

	if orig.Idmap[0].Hostid != 500 || orig.Idmap[0].Nsid != 0 || orig.Idmap[0].Maprange != 10 {
		t.Error(fmt.Errorf("bad range: %v", orig.Idmap[0]))
		return
	}

	if orig.Idmap[1].Hostid != 1010 || orig.Idmap[1].Nsid != 10 || orig.Idmap[1].Maprange != 990 {
		t.Error(fmt.Errorf("bad range: %v", orig.Idmap[1]))
		return
	}

	if len(orig.Idmap) != 2 {
		t.Error("too many idmap entries")
		return
	}
}

func TestIdmapSetAddSafe_upper(t *testing.T) {
	orig := IdmapSet{Idmap: []IdmapEntry{IdmapEntry{Isuid: true, Hostid: 1000, Nsid: 0, Maprange: 1000}}}

	if err := orig.AddSafe(IdmapEntry{Isuid: true, Hostid: 500, Nsid: 995, Maprange: 10}); err != nil {
		t.Error(err)
		return
	}

	if orig.Idmap[0].Hostid != 1000 || orig.Idmap[0].Nsid != 0 || orig.Idmap[0].Maprange != 995 {
		t.Error(fmt.Errorf("bad range: %v", orig.Idmap[0]))
		return
	}

	if orig.Idmap[1].Hostid != 500 || orig.Idmap[1].Nsid != 995 || orig.Idmap[1].Maprange != 10 {
		t.Error(fmt.Errorf("bad range: %v", orig.Idmap[1]))
		return
	}

	if len(orig.Idmap) != 2 {
		t.Error("too many idmap entries")
		return
	}
}
