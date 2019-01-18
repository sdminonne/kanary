package utils

import (
	"testing"

	corev1 "k8s.io/api/core/v1"

	kanaryv1alpha1 "github.com/amadeusitgroup/kanary/pkg/apis/kanary/v1alpha1"
)

func TestIsKanaryDeploymentFailed(t *testing.T) {
	type args struct {
		status *kanaryv1alpha1.KanaryDeploymentStatus
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "failed",
			args: args{
				status: &kanaryv1alpha1.KanaryDeploymentStatus{
					Conditions: []kanaryv1alpha1.KanaryDeploymentCondition{
						{
							Type:   kanaryv1alpha1.FailedKanaryDeploymentConditionType,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "not failed",
			args: args{
				status: &kanaryv1alpha1.KanaryDeploymentStatus{},
			},
			want: false,
		},
		{
			name: "not failed, conditionFalse",
			args: args{
				status: &kanaryv1alpha1.KanaryDeploymentStatus{
					Conditions: []kanaryv1alpha1.KanaryDeploymentCondition{
						{
							Type:   kanaryv1alpha1.FailedKanaryDeploymentConditionType,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKanaryDeploymentFailed(tt.args.status); got != tt.want {
				t.Errorf("IsKanaryDeploymentFailed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsKanaryDeploymentSucceeded(t *testing.T) {
	type args struct {
		status *kanaryv1alpha1.KanaryDeploymentStatus
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "succeed",
			args: args{
				status: &kanaryv1alpha1.KanaryDeploymentStatus{
					Conditions: []kanaryv1alpha1.KanaryDeploymentCondition{
						{
							Type:   kanaryv1alpha1.SucceededKanaryDeploymentConditionType,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "not succeed",
			args: args{
				status: &kanaryv1alpha1.KanaryDeploymentStatus{},
			},
			want: false,
		},
		{
			name: "not succeed, conditionFalse",
			args: args{
				status: &kanaryv1alpha1.KanaryDeploymentStatus{
					Conditions: []kanaryv1alpha1.KanaryDeploymentCondition{
						{
							Type:   kanaryv1alpha1.SucceededKanaryDeploymentConditionType,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsKanaryDeploymentSucceeded(tt.args.status); got != tt.want {
				t.Errorf("IsKanaryDeploymentSucceeded() = %v, want %v", got, tt.want)
			}
		})
	}
}