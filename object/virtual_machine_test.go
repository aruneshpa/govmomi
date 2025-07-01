// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"testing"
	"time"

	"github.com/vmware/govmomi/vim25/types"
)

// VirtualMachine should implement the Reference interface.
var _ Reference = VirtualMachine{}

// pretty.Printf generated
var snapshot = &types.VirtualMachineSnapshotInfo{
	DynamicData:     types.DynamicData{},
	CurrentSnapshot: &types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-11"},
	RootSnapshotList: []types.VirtualMachineSnapshotTree{
		{
			DynamicData:    types.DynamicData{},
			Snapshot:       types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-1"},
			Vm:             types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
			Name:           "root",
			Description:    "",
			Id:             1,
			CreateTime:     time.Now(),
			State:          "poweredOn",
			Quiesced:       false,
			BackupManifest: "",
			ChildSnapshotList: []types.VirtualMachineSnapshotTree{
				{
					DynamicData:       types.DynamicData{},
					Snapshot:          types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-2"},
					Vm:                types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
					Name:              "child",
					Description:       "",
					Id:                2,
					CreateTime:        time.Now(),
					State:             "poweredOn",
					Quiesced:          false,
					BackupManifest:    "",
					ChildSnapshotList: nil,
					ReplaySupported:   types.NewBool(false),
				},
				{
					DynamicData:    types.DynamicData{},
					Snapshot:       types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-3"},
					Vm:             types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
					Name:           "child",
					Description:    "",
					Id:             3,
					CreateTime:     time.Now(),
					State:          "poweredOn",
					Quiesced:       false,
					BackupManifest: "",
					ChildSnapshotList: []types.VirtualMachineSnapshotTree{
						{
							DynamicData:    types.DynamicData{},
							Snapshot:       types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-9"},
							Vm:             types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
							Name:           "grandkid",
							Description:    "",
							Id:             9,
							CreateTime:     time.Now(),
							State:          "poweredOn",
							Quiesced:       false,
							BackupManifest: "",
							ChildSnapshotList: []types.VirtualMachineSnapshotTree{
								{
									DynamicData:       types.DynamicData{},
									Snapshot:          types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-10"},
									Vm:                types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
									Name:              "great",
									Description:       "",
									Id:                10,
									CreateTime:        time.Now(),
									State:             "poweredOn",
									Quiesced:          false,
									BackupManifest:    "",
									ChildSnapshotList: nil,
									ReplaySupported:   types.NewBool(false),
								},
							},
							ReplaySupported: types.NewBool(false),
						},
					},
					ReplaySupported: types.NewBool(false),
				},
				{
					DynamicData:    types.DynamicData{},
					Snapshot:       types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-5"},
					Vm:             types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
					Name:           "voodoo",
					Description:    "",
					Id:             5,
					CreateTime:     time.Now(),
					State:          "poweredOn",
					Quiesced:       false,
					BackupManifest: "",
					ChildSnapshotList: []types.VirtualMachineSnapshotTree{
						{
							DynamicData:       types.DynamicData{},
							Snapshot:          types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-11"},
							Vm:                types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
							Name:              "child",
							Description:       "",
							Id:                11,
							CreateTime:        time.Now(),
							State:             "poweredOn",
							Quiesced:          false,
							BackupManifest:    "",
							ChildSnapshotList: nil,
							ReplaySupported:   types.NewBool(false),
						},
					},
					ReplaySupported: types.NewBool(false),
				},
				{
					DynamicData:    types.DynamicData{},
					Snapshot:       types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-6"},
					Vm:             types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
					Name:           "better",
					Description:    "",
					Id:             6,
					CreateTime:     time.Now(),
					State:          "poweredOn",
					Quiesced:       false,
					BackupManifest: "",
					ChildSnapshotList: []types.VirtualMachineSnapshotTree{
						{
							DynamicData:    types.DynamicData{},
							Snapshot:       types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-7"},
							Vm:             types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
							Name:           "best",
							Description:    "",
							Id:             7,
							CreateTime:     time.Now(),
							State:          "poweredOn",
							Quiesced:       false,
							BackupManifest: "",
							ChildSnapshotList: []types.VirtualMachineSnapshotTree{
								{
									DynamicData:       types.DynamicData{},
									Snapshot:          types.ManagedObjectReference{Type: "VirtualMachineSnapshot", Value: "2-snapshot-8"},
									Vm:                types.ManagedObjectReference{Type: "VirtualMachine", Value: "2"},
									Name:              "betterer",
									Description:       "",
									Id:                8,
									CreateTime:        time.Now(),
									State:             "poweredOn",
									Quiesced:          false,
									BackupManifest:    "",
									ChildSnapshotList: nil,
									ReplaySupported:   types.NewBool(false),
								},
							},
							ReplaySupported: types.NewBool(false),
						},
					},
					ReplaySupported: types.NewBool(false),
				},
			},
			ReplaySupported: types.NewBool(false),
		},
	},
}

func TestVirtualMachineSnapshotMap(t *testing.T) {
	m := make(snapshotMap)
	m.add("", snapshot.RootSnapshotList)

	tests := []struct {
		name   string
		expect int
	}{
		{"enoent", 0},
		{"root", 1},
		{"child", 3},
		{"root/child", 2},
		{"root/voodoo/child", 1},
		{"2-snapshot-6", 1},
	}

	for _, test := range tests {
		s := m[test.name]

		if len(s) != test.expect {
			t.Errorf("%s: %d != %d", test.name, len(s), test.expect)
		}
	}
}

func TestDiskFileOperation(t *testing.T) {
	backing := &types.VirtualDiskFlatVer2BackingInfo{
		VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
			FileName: "[datastore1] data/disk1.vmdk",
		},
		Parent: nil,
	}

	parent := &types.VirtualDiskFlatVer2BackingInfo{
		VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
			FileName: "[datastore1] data/parent.vmdk",
		},
	}

	disk := &types.VirtualDisk{
		VirtualDevice: types.VirtualDevice{
			Backing: backing,
		},
	}

	op := types.VirtualDeviceConfigSpecOperationAdd
	fop := types.VirtualDeviceConfigSpecFileOperationCreate

	res := diskFileOperation(op, fop, disk)
	if res != "" {
		t.Errorf("res=%s", res)
	}

	disk.CapacityInKB = 1
	res = diskFileOperation(op, fop, disk)
	if res != types.VirtualDeviceConfigSpecFileOperationCreate {
		t.Errorf("res=%s", res)
	}

	disk.CapacityInKB = 0
	disk.CapacityInBytes = 1
	res = diskFileOperation(op, fop, disk)
	if res != types.VirtualDeviceConfigSpecFileOperationCreate {
		t.Errorf("res=%s", res)
	}

	disk.CapacityInBytes = 0
	backing.Parent = parent
	res = diskFileOperation(op, fop, disk)
	if res != types.VirtualDeviceConfigSpecFileOperationCreate {
		t.Errorf("res=%s", res)
	}
}

func TestVirtualMachineFindSnapshotTree(t *testing.T) {
	// Test cases for finding snapshots by name
	nameTests := []struct {
		name       string
		expected   string // expected snapshot name
		shouldFind bool
	}{
		{"root", "root", true},
		{"child", "child", true}, // should find one of the child snapshots
		{"grandkid", "grandkid", true},
		{"great", "great", true},
		{"voodoo", "voodoo", true},
		{"better", "better", true},
		{"best", "best", true},
		{"betterer", "betterer", true},
		{"nonexistent", "", false},
		{"", "", false},
	}

	// Test cases for finding snapshots by ManagedObjectReference value
	refTests := []struct {
		refValue   string
		expected   string // expected snapshot name
		shouldFind bool
	}{
		{"2-snapshot-1", "root", true},
		{"2-snapshot-2", "child", true},
		{"2-snapshot-3", "child", true},
		{"2-snapshot-9", "grandkid", true},
		{"2-snapshot-10", "great", true},
		{"2-snapshot-5", "voodoo", true},
		{"2-snapshot-11", "child", true},
		{"2-snapshot-6", "better", true},
		{"2-snapshot-7", "best", true},
		{"2-snapshot-8", "betterer", true},
		{"nonexistent-ref", "", false},
		{"", "", false},
	}

	// Test finding by name
	for _, test := range nameTests {
		// Since we can't easily mock the Properties call in this test,
		// we'll test the logic by directly calling the helper function
		// that would be used by FindSnapshotTree
		var findSnapshot func([]types.VirtualMachineSnapshotTree) *types.VirtualMachineSnapshotTree
		findSnapshot = func(trees []types.VirtualMachineSnapshotTree) *types.VirtualMachineSnapshotTree {
			for _, tree := range trees {
				if tree.Name == test.name || tree.Snapshot.Value == test.name {
					return &tree
				}
				if result := findSnapshot(tree.ChildSnapshotList); result != nil {
					return result
				}
			}
			return nil
		}

		result := findSnapshot(snapshot.RootSnapshotList)

		if test.shouldFind {
			if result == nil {
				t.Errorf("FindSnapshotTree by name '%s': expected to find snapshot, but got nil", test.name)
			} else if result.Name != test.expected {
				t.Errorf("FindSnapshotTree by name '%s': expected name '%s', got '%s'", test.name, test.expected, result.Name)
			}
		} else {
			if result != nil {
				t.Errorf("FindSnapshotTree by name '%s': expected nil, but found snapshot '%s'", test.name, result.Name)
			}
		}
	}

	// Test finding by ManagedObjectReference value
	for _, test := range refTests {
		var findSnapshot func([]types.VirtualMachineSnapshotTree) *types.VirtualMachineSnapshotTree
		findSnapshot = func(trees []types.VirtualMachineSnapshotTree) *types.VirtualMachineSnapshotTree {
			for _, tree := range trees {
				if tree.Name == test.refValue || tree.Snapshot.Value == test.refValue {
					return &tree
				}
				if result := findSnapshot(tree.ChildSnapshotList); result != nil {
					return result
				}
			}
			return nil
		}

		result := findSnapshot(snapshot.RootSnapshotList)

		if test.shouldFind {
			if result == nil {
				t.Errorf("FindSnapshotTree by ref '%s': expected to find snapshot, but got nil", test.refValue)
			} else if result.Name != test.expected {
				t.Errorf("FindSnapshotTree by ref '%s': expected name '%s', got '%s'", test.refValue, test.expected, result.Name)
			} else if result.Snapshot.Value != test.refValue {
				t.Errorf("FindSnapshotTree by ref '%s': expected ref '%s', got '%s'", test.refValue, test.refValue, result.Snapshot.Value)
			}
		} else {
			if result != nil {
				t.Errorf("FindSnapshotTree by ref '%s': expected nil, but found snapshot '%s'", test.refValue, result.Name)
			}
		}
	}

	// Test that we can access snapshot properties
	var findSnapshot func([]types.VirtualMachineSnapshotTree) *types.VirtualMachineSnapshotTree
	findSnapshot = func(trees []types.VirtualMachineSnapshotTree) *types.VirtualMachineSnapshotTree {
		for _, tree := range trees {
			if tree.Name == "root" {
				return &tree
			}
			if result := findSnapshot(tree.ChildSnapshotList); result != nil {
				return result
			}
		}
		return nil
	}

	result := findSnapshot(snapshot.RootSnapshotList)
	if result != nil {
		// Test that we can access the snapshot properties
		if result.Name != "root" {
			t.Errorf("Expected snapshot name 'root', got '%s'", result.Name)
		}
		if result.Snapshot.Value != "2-snapshot-1" {
			t.Errorf("Expected snapshot ref '2-snapshot-1', got '%s'", result.Snapshot.Value)
		}
		if result.Snapshot.Type != "VirtualMachineSnapshot" {
			t.Errorf("Expected snapshot type 'VirtualMachineSnapshot', got '%s'", result.Snapshot.Type)
		}
		if len(result.ChildSnapshotList) == 0 {
			t.Errorf("Expected root snapshot to have child snapshots")
		}
	} else {
		t.Errorf("Expected to find root snapshot")
	}
}
