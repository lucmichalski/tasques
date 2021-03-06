package config

import "time"

type Tasques struct {
	Server App `json:"server" mapstructure:"server"`
}

type TopLevel struct {
	Tasques Tasques `json:"tasques" mapstructure:"tasques"`
}

type App struct {
	BindAddress     string              `json:"bind_address" mapstructure:"bind_address"`
	ShutdownTimeout time.Duration       `json:"shutdown_timeout" mapstructure:"shutdown_timeout"`
	Elasticsearch   ElasticsearchClient `json:"elasticsearch" mapstructure:"elasticsearch"`
	KibanaClient    *KibanaClient       `json:"kibana,omitempty" mapstructure:"kibana"`
	ApmClient       *ApmClient          `json:"apm,omitempty" mapstructure:"apm"`
	Auth            *Auth               `json:"auth,omitempty" mapstructure:"auth"`
	Logging         *Logging            `json:"logging,omitempty" mapstructure:"logging"`
	Tasks           Tasks               `json:"tasks" mapstructure:"tasks"`
	Recurring       Recurring           `json:"recurring" mapstructure:"recurring"`
	LifecycleSetup  LifecycleSetup      `json:"lifecycle_setup" mapstructure:"lifecycle_setup"`
}

type Logging struct {
	Json  *bool   `json:"json,omitempty" mapstructure:"json"`
	File  *string `json:"file,omitempty" mapstructure:"file"`
	Level *string `json:"level,omitempty" mapstructure:"level"`
}

type KibanaClient struct {
	Address string         `json:"address" mapstructure:"address"`
	User    *BasicAuthUser `json:"user,omitempty" mapstructure:"user"`
}

type ElasticsearchClient struct {
	Addresses []string       `json:"addresses" mapstructure:"addresses"`
	User      *BasicAuthUser `json:"user,omitempty" mapstructure:"user"`
}

type ApmClient struct {
	Address     *string `json:"address,omitempty" mapstructure:"address"`
	SecretToken *string `json:"secret_token,omitempty" mapstructure:"secret_token"`
}

type Tasks struct {
	Defaults TasksDefaults `json:"defaults" mapstructure:"defaults"`
	Queues   Queues        `json:"queues" mapstructure:"queues"`
}

type Queues struct {
	LastActivityTrackerMaxSize uint          `json:"last_activity_tracker_max_size" mapstructure:"last_activity_tracker_max_size"`
	RefreshIfLastTouchedOver   time.Duration `json:"refresh_if_last_touched_over" mapstructure:"refresh_if_last_touched_over"`
	RefreshIfLastTouchedUnder  time.Duration `json:"refresh_if_last_touched_under" mapstructure:"refresh_if_last_touched_under"`
}

type TasksDefaults struct {
	BlockFor                    time.Duration `json:"block_for" mapstructure:"block_for"`
	BlockForRetryMinWait        time.Duration `json:"block_for_retry_min_wait" mapstructure:"block_for_retry_min_wait"`
	BlockForRetryMaxRetries     uint          `json:"block_for_retry_max_retries" mapstructure:"block_for_retry_max_retries"`
	WorkerProcessingTimeout     time.Duration `json:"worker_processing_timeout" mapstructure:"worker_processing_timeout"`
	ClaimAmount                 uint          `json:"claim_amount" mapstructure:"claim_amount"`
	ClaimAmountSearchMultiplier uint          `json:"claim_amount_search_multiplier" mapstructure:"claim_amount_search_multiplier"`
	RetryTimes                  uint          `json:"retry_times" mapstructure:"retry_times"`
	VersionConflictRetryTimes   uint          `json:"version_conflict_retry_times" mapstructure:"version_conflict_retry_times"`
}

type TimedOutTasksReaper struct {
	ScrollSize  uint          `json:"scroll_size" mapstructure:"scroll_size"`
	ScrollTtl   time.Duration `json:"scroll_ttl" mapstructure:"scroll_ttl"`
	RunInterval time.Duration `json:"run_interval" mapstructure:"run_interval"`
}

type TasksArchiver struct {
	ScrollSize       uint          `json:"scroll_size" mapstructure:"scroll_size"`
	ScrollTtl        time.Duration `json:"scroll_ttl" mapstructure:"scroll_ttl"`
	ArchiveOlderThan time.Duration `json:"archive_older_than" mapstructure:"archive_older_than"`
	RunInterval      time.Duration `json:"run_interval" mapstructure:"run_interval"`
}

type RecurringTasks struct {
	ScrollSize             uint          `json:"scroll_size" mapstructure:"scroll_size"`
	ScrollTtl              time.Duration `json:"scroll_ttl" mapstructure:"scroll_ttl"`
	SyncRunInterval        time.Duration `json:"sync_run_interval" mapstructure:"sync_run_interval"`
	EnforceSyncRunInterval time.Duration `json:"enforce_sync_run_interval" mapstructure:"enforce_sync_run_interval"`
}

type Auth struct {
	BasicAuth []BasicAuthUser `json:"basic_auth" mapstructure:"basic_auth"`
}

type BasicAuthUser struct {
	Name     string `json:"name" mapstructure:"name"`
	Password string `json:"password" mapstructure:"password"`
}

type Recurring struct {
	LeaderLock          LeaderLock          `json:"leader_lock" mapstructure:"leader_lock"`
	TimedOutTasksReaper TimedOutTasksReaper `json:"timed_out_tasks_reaper" mapstructure:"timed_out_tasks_reaper"`
	TasksArchiver       TasksArchiver       `json:"tasks_archiver" mapstructure:"tasks_archiver"`
	RecurringTasks      RecurringTasks      `json:"recurring_tasks" mapstructure:"recurring_tasks"`
}

type LeaderLock struct {
	RunInterval     time.Duration `json:"run_interval" mapstructure:"run_interval"`
	LeaderLockLease time.Duration `json:"leader_lock_lease" mapstructure:"leader_lock_lease"`
}

type LifecycleSetup struct {
	ArchivedTasks LifecycleSettings `json:"archived_tasks" mapstructure:"archived_tasks"`
}

type LifecycleSettings struct {
	Enabled      bool                   `json:"enabled" mapstructure:"enabled"`
	CustomPolicy *CustomLifecyclePolicy `json:"custom_policy,omitempty" mapstructure:"custom_policy"`
}

type CustomLifecyclePolicy struct {
	Name   string                 `json:"name" mapstructure:"name"`
	Policy map[string]interface{} `json:"policy" mapstructure:"policy"`
}
