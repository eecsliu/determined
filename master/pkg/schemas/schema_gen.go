// This is a generated file.  Editing it will make you sad.

package schemas

import (
	"encoding/json"
)

var (
	textBindMountV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/bind-mount.json",
    "title": "BindMount",
    "additionalProperties": false,
    "required": [
        "host_path",
        "container_path"
    ],
    "type": "object",
    "properties": {
        "host_path": {
            "type": "string",
            "checks": {
                "host_path must be an absolute path": {
                    "pattern": "^/"
                }
            }
        },
        "container_path": {
            "type": "string",
            "checks": {
                "container_path must not be \".\"": {
                    "not": {
                        "const": "."
                    }
                }
            }
        },
        "read_only": {
            "type": [
                "boolean",
                "null"
            ],
            "default": false
        },
        "propagation": {
            "type": [
                "string",
                "null"
            ],
            "default": "rprivate"
        }
    }
}
`)
	textCheckDataLayerCacheV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/check-data-layer-cache.json",
    "title": "CheckDataLayerCache",
    "checks": {
        "local_cache_container_path must be specified if local_cache_host_path is set": {
            "not": {
                "required": [
                    "local_cache_host_path"
                ],
                "properties": {
                    "local_cache_container_path": {
                        "type": "null"
                    },
                    "local_cache_host_path": {
                        "type": "string"
                    }
                }
            }
        },
        "local_cache_host_path must be specified if local_cache_container_path is set": {
            "not": {
                "required": [
                    "local_cache_container_path"
                ],
                "properties": {
                    "local_cache_container_path": {
                        "type": "string"
                    },
                    "local_cache_host_path": {
                        "type": "null"
                    }
                }
            }
        }
    }
}
`)
	textCheckEpochNotUsedV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/check-epoch-not-used.json",
    "title": "CheckEpochNotUsed",
    "additionalProperties": {
        "$ref": "http://determined.ai/schemas/expconf/v0/check-epoch-not-used.json"
    },
    "items": {
        "$ref": "http://determined.ai/schemas/expconf/v0/check-epoch-not-used.json"
    },
    "checks": {
        "must specify the top-level records_per_epoch when this field is in terms of epochs": {
            "properties": {
                "epochs": {
                    "not": {
                        "type": "number"
                    }
                }
            }
        }
    }
}
`)
	textCheckGlobalBatchSizeV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/check-global-batch-size.json",
    "title": "CheckGlobalBatchSize",
    "union": {
        "defaultMessage": "is neither a positive integer nor an int hyperparameter",
        "items": [
            {
                "unionKey": "const:type=int",
                "allOf": [
                    {
                        "$ref": "http://determined.ai/schemas/expconf/v0/hyperparameter-int.json"
                    },
                    {
                        "properties": {
                            "minval": {
                                "type": "number",
                                "minimum": 1
                            }
                        }
                    }
                ]
            },
            {
                "unionKey": "const:type=const",
                "allOf": [
                    {
                        "$ref": "http://determined.ai/schemas/expconf/v0/hyperparameter-const.json"
                    },
                    {
                        "properties": {
                            "val": {
                                "type": "number",
                                "minimum": 1
                            }
                        }
                    }
                ]
            },
            {
                "unionKey": "const:type=categorical",
                "allOf": [
                    {
                        "$ref": "http://determined.ai/schemas/expconf/v0/hyperparameter-categorical.json"
                    },
                    {
                        "properties": {
                            "vals": {
                                "type": "array",
                                "items": {
                                    "type": "integer",
                                    "minimum": 1
                                }
                            }
                        }
                    }
                ]
            },
            {
                "unionKey": "never",
                "type": "integer",
                "minimum": 1
            }
        ]
    }
}
`)
	textCheckGridHyperparameterV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/check-grid-hyperparameter.json",
    "title": "CheckGridHyperparameter",
    "union": {
        "items": [
            {
                "unionKey": "type:array",
                "type": "array",
                "items": {
                    "$ref": "http://determined.ai/schemas/expconf/v0/check-grid-hyperparameter.json"
                }
            },
            {
                "unionKey": "not:hasattr:type",
                "type": "object",
                "properties": {
                    "type": false
                },
                "additionalProperties": {
                    "$ref": "http://determined.ai/schemas/expconf/v0/check-grid-hyperparameter.json"
                }
            },
            {
                "unionKey": "never",
                "not": {
                    "type": [
                        "object",
                        "array"
                    ]
                }
            },
            {
                "unionKey": "hasattr:type",
                "type": "object",
                "required": [
                    "type"
                ],
                "properties": {
                    "type": {
                        "type": "string"
                    }
                },
                "checks": {
                    "grid search is in use but count was not provided": {
                        "conditional": {
                            "$comment": "unless type is not double/log/int, expect non-null count",
                            "unless": {
                                "not": {
                                    "properties": {
                                        "type": {
                                            "enum": [
                                                "double",
                                                "log",
                                                "int"
                                            ]
                                        }
                                    }
                                }
                            },
                            "enforce": {
                                "not": {
                                    "properties": {
                                        "count": {
                                            "type": "null"
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        ]
    }
}
`)
	textCheckPositiveLengthV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/check-positive-length.json",
    "title": "CheckPositiveLength",
    "allOf": [
        {
            "$ref": "http://determined.ai/schemas/expconf/v0/length.json"
        },
        {
            "additionalProperties": {
                "type": "integer",
                "minimum": 1
            }
        }
    ]
}
`)
	textCheckpointStorageConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/checkpoint-storage.json",
    "title": "CheckpointStorageConfig",
    "union": {
        "defaultMessage": "is not an object where object[\"type\"] is one of 'shared_fs', 'hdfs', 's3', or 'gcs'",
        "items": [
            {
                "unionKey": "const:type=shared_fs",
                "$ref": "http://determined.ai/schemas/expconf/v0/shared-fs.json"
            },
            {
                "unionKey": "const:type=hdfs",
                "$ref": "http://determined.ai/schemas/expconf/v0/hdfs.json"
            },
            {
                "unionKey": "const:type=s3",
                "$ref": "http://determined.ai/schemas/expconf/v0/s3.json"
            },
            {
                "unionKey": "const:type=gcs",
                "$ref": "http://determined.ai/schemas/expconf/v0/gcs.json"
            }
        ]
    }
}
`)
	textGCSDataLayerConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/data-layer-gcs.json",
    "title": "GCSDataLayerConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "type",
        "bucket",
        "bucket_directory_path"
    ],
    "properties": {
        "type": {
            "const": "gcs"
        },
        "bucket": {
            "type": "string"
        },
        "bucket_directory_path": {
            "type": "string"
        },
        "local_cache_host_path": {
            "type": [
                "string",
                "null"
            ],
            "checks": {
                "local_cache_host_path must be an absolute path": {
                    "pattern": "^/"
                }
            },
            "default": null
        },
        "local_cache_container_path": {
            "type": [
                "string",
                "null"
            ],
            "checks": {
                "local_cache_container_path must be an absolute path": {
                    "pattern": "^/"
                }
            },
            "default": null
        }
    },
    "allOf": [
        {
            "$ref": "http://determined.ai/schemas/expconf/v0/check-data-layer-cache.json"
        }
    ]
}
`)
	textS3DataLayerConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/data-layer-s3.json",
    "title": "S3DataLayerConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "type",
        "bucket",
        "bucket_directory_path"
    ],
    "properties": {
        "type": {
            "const": "s3"
        },
        "bucket": {
            "type": "string"
        },
        "bucket_directory_path": {
            "type": "string"
        },
        "local_cache_host_path": {
            "type": [
                "string",
                "null"
            ],
            "checks": {
                "local_cache_host_path must be an absolute path": {
                    "pattern": "^/"
                }
            },
            "default": null
        },
        "local_cache_container_path": {
            "type": [
                "string",
                "null"
            ],
            "checks": {
                "local_cache_container_path must be an absolute path": {
                    "pattern": "^/"
                }
            },
            "default": null
        },
        "access_key": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        },
        "secret_key": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        },
        "endpoint_url": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        }
    },
    "allOf": [
        {
            "$ref": "http://determined.ai/schemas/expconf/v0/check-data-layer-cache.json"
        }
    ]
}
`)
	textSharedFSDataLayerConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/data-layer-shared-fs.json",
    "title": "SharedFSDataLayerConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "type"
    ],
    "properties": {
        "type": {
            "const": "shared_fs"
        },
        "host_storage_path": {
            "type": [
                "string",
                "null"
            ],
            "checks": {
                "host_storage_path must be an absolute path": {
                    "pattern": "^/"
                }
            },
            "default": null
        },
        "container_storage_path": {
            "type": [
                "string",
                "null"
            ],
            "checks": {
                "container_storage_path must be an absolute path": {
                    "pattern": "^/"
                }
            },
            "default": null
        }
    }
}
`)
	textDataLayerConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/data-layer.json",
    "title": "DataLayerConfig",
    "union": {
        "defaultMessage": "is not an object where object[\"type\"] is one of 'shared_fs', 's3', or 'gcs'",
        "items": [
            {
                "unionKey": "const:type=shared_fs",
                "$ref": "http://determined.ai/schemas/expconf/v0/data-layer-shared-fs.json"
            },
            {
                "unionKey": "const:type=gcs",
                "$ref": "http://determined.ai/schemas/expconf/v0/data-layer-gcs.json"
            },
            {
                "unionKey": "const:type=s3",
                "$ref": "http://determined.ai/schemas/expconf/v0/data-layer-s3.json"
            }
        ]
    }
}
`)
	textEnvironmentImageMapV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/environment-image-map.json",
    "title": "EnvironmentImageMap",
    "type": "object",
    "additionalProperties": false,
    "required": [],
    "properties": {
        "cpu": {
            "type": [
                "string",
                "null"
            ],
            "default": ""
        },
        "gpu": {
            "type": [
                "string",
                "null"
            ],
            "default": ""
        }
    }
}
`)
	textEnvironmentImageV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/environment-image.json",
    "title": "EnvironmentImage",
    "union": {
        "defaultMessage": "is neither a string nor a map of cpu/gpu to strings",
        "items": [
            {
                "unionKey": "never",
                "$ref": "http://determined.ai/schemas/expconf/v0/environment-image-map.json"
            },
            {
                "unionKey": "never",
                "type": "string"
            }
        ]
    }
}
`)
	textEnvironmentVariablesMapV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/environment-variables-map.json",
    "title": "EnvironmentVariablesMap",
    "type": "object",
    "additionalProperties": false,
    "properties": {
        "cpu": {
            "type": [
                "array",
                "null"
            ],
            "default": [],
            "items": {
                "type": "string"
            }
        },
        "gpu": {
            "type": [
                "array",
                "null"
            ],
            "default": [],
            "items": {
                "type": "string"
            }
        }
    }
}
`)
	textEnvironmentVariablesV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/environment-variables.json",
    "title": "EnvironmentVariables",
    "union": {
        "defaultMessage": "is neither a list of strings nor a map of cpu/gpu to lists of strings",
        "items": [
            {
                "unionKey": "never",
                "$ref": "http://determined.ai/schemas/expconf/v0/environment-variables-map.json"
            },
            {
                "unionKey": "never",
                "type": "array",
                "items": {
                    "type": "string"
                }
            }
        ]
    }
}
`)
	textEnvironmentConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/environment.json",
    "title": "EnvironmentConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [],
    "properties": {
        "image": {
            "type": [
                "object",
                "null"
            ],
            "default": {},
            "optionalRef": "http://determined.ai/schemas/expconf/v0/environment-image.json"
        },
        "environment_variables": {
            "type": [
                "object",
                "array",
                "null"
            ],
            "default": [],
            "optionalRef": "http://determined.ai/schemas/expconf/v0/environment-variables.json"
        },
        "ports": {
            "type": [
                "object",
                "null"
            ],
            "default": {},
            "additionalProperties": {
                "type": "integer"
            }
        },
        "force_pull_image": {
            "type": [
                "boolean",
                "null"
            ],
            "default": false
        },
        "pod_spec": {
            "type": [
                "object",
                "null"
            ],
            "default": null,
            "disallowProperties": {
                "name": "pod Name is not a configurable option",
                "name_space": "pod NameSpace is not a configurable option"
            },
            "properties": {
                "spec": {
                    "type": [
                        "object",
                        "null"
                    ],
                    "default": null,
                    "properties": {
                        "containers": {
                            "type": [
                                "array",
                                "null"
                            ],
                            "default": null,
                            "items": {
                                "type": "object",
                                "disallowProperties": {
                                    "image": "container Image is not configurable, set it in the experiment config",
                                    "command": "container Command is not configurable",
                                    "args": "container Args are not configurable",
                                    "working_dir": "container WorkingDir is not configurable",
                                    "ports": "container Ports are not configurable",
                                    "env_from": "container EnvFrom is not configurable",
                                    "env": "container Env is not configurable, set it in the experiment config",
                                    "liveness_probe": "container LivenessProbe is not configurable",
                                    "readiness_probe": "container ReadinessProbe is not configurable",
                                    "startup_probe": "container StartupProbe is not configurable",
                                    "lifecycle": "container Lifecycle is not configurable",
                                    "termination_message_path": "container TerminationMessagePath is not configurable",
                                    "termination_message_policy": "container TerminationMessagePolicy is not configurable",
                                    "image_pull_policy": "container ImagePullPolicy is not configurable, set it in the experiment config",
                                    "security_context": "container SecurityContext is not configurable, set it in the experiment config"
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
`)
	textExperimentConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/experiment.json",
    "title": "ExperimentConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "hyperparameters",
        "searcher"
    ],
    "eventuallyRequired": [
        "checkpoint_storage"
    ],
    "properties": {
        "bind_mounts": {
            "type": [
                "array",
                "null"
            ],
            "items": {
                "$ref": "http://determined.ai/schemas/expconf/v0/bind-mount.json"
            },
            "default": []
        },
        "checkpoint_policy": {
            "enum": [
                null,
                "best",
                "all",
                "none"
            ],
            "default": "best"
        },
        "checkpoint_storage": {
            "type": [
                "object",
                "null"
            ],
            "default": null,
            "optionalRef": "http://determined.ai/schemas/expconf/v0/checkpoint-storage.json"
        },
        "data": {
            "type": [
                "object",
                "null"
            ],
            "default": {}
        },
        "data_layer": {
            "type": [
                "object",
                "null"
            ],
            "default": {
                "type": "shared_fs"
            },
            "optionalRef": "http://determined.ai/schemas/expconf/v0/data-layer.json"
        },
        "debug": {
            "type": [
                "boolean",
                "null"
            ],
            "default": false
        },
        "description": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        },
        "entrypoint": {
            "type": [
                "string",
                "null"
            ],
            "checks": {
                "entrypoint must be of the form \"module.submodule:ClassName\"": {
                    "pattern": "^[a-zA-Z0-9_.]+:[a-zA-Z0-9_]+$"
                }
            },
            "default": null
        },
        "environment": {
            "type": [
                "object",
                "null"
            ],
            "default": {},
            "optionalRef": "http://determined.ai/schemas/expconf/v0/environment.json"
        },
        "hyperparameters": {
            "$ref": "http://determined.ai/schemas/expconf/v0/hyperparameters.json"
        },
        "internal": {
            "type": [
                "object",
                "null"
            ],
            "default": null,
            "optionalRef": "http://determined.ai/schemas/expconf/v0/internal.json"
        },
        "labels": {
            "type": [
                "array",
                "null"
            ],
            "default": [],
            "items": {
                "type": "string"
            }
        },
        "max_restarts": {
            "type": [
                "integer",
                "null"
            ],
            "minimum": 0,
            "default": 5
        },
        "min_checkpoint_period": {
            "type": [
                "object",
                "null"
            ],
            "default": {
                "batches": 0
            },
            "optionalRef": "http://determined.ai/schemas/expconf/v0/length.json"
        },
        "min_validation_period": {
            "type": [
                "object",
                "null"
            ],
            "default": {
                "batches": 0
            },
            "optionalRef": "http://determined.ai/schemas/expconf/v0/length.json"
        },
        "optimizations": {
            "type": [
                "object",
                "null"
            ],
            "default": {},
            "optionalRef": "http://determined.ai/schemas/expconf/v0/optimizations.json"
        },
        "perform_initial_validation": {
            "type": [
                "boolean",
                "null"
            ],
            "default": false
        },
        "records_per_epoch": {
            "type": [
                "integer",
                "null"
            ],
            "default": 0
        },
        "reproducibility": {
            "type": [
                "object",
                "null"
            ],
            "default": {},
            "optionalRef": "http://determined.ai/schemas/expconf/v0/reproducibility.json"
        },
        "resources": {
            "type": [
                "object",
                "null"
            ],
            "default": {},
            "optionalRef": "http://determined.ai/schemas/expconf/v0/resources.json"
        },
        "scheduling_unit": {
            "type": [
                "integer",
                "null"
            ],
            "minimum": 1,
            "default": 100
        },
        "searcher": {
            "$ref": "http://determined.ai/schemas/expconf/v0/searcher.json"
        },
        "security": {
            "type": [
                "object",
                "null"
            ],
            "default": null,
            "optionalRef": "http://determined.ai/schemas/expconf/v0/security.json"
        },
        "tensorboard_storage": {
            "type": [
                "object",
                "null"
            ],
            "default": null,
            "optionalRef": "http://determined.ai/schemas/expconf/v0/tensorboard-storage.json"
        }
    },
    "allOf": [
        {
            "conditional": {
                "$comment": "when grid search is in use, expect hp counts",
                "when": {
                    "properties": {
                        "searcher": {
                            "properties": {
                                "name": {
                                    "const": "grid"
                                }
                            }
                        }
                    }
                },
                "enforce": {
                    "properties": {
                        "hyperparameters": {
                            "additionalProperties": {
                                "$ref": "http://determined.ai/schemas/expconf/v0/check-grid-hyperparameter.json"
                            }
                        }
                    }
                }
            }
        },
        {
            "conditional": {
                "$comment": "when records per epoch not set, forbid epoch lengths",
                "when": {
                    "properties": {
                        "records_per_epoch": {
                            "maximum": 0
                        }
                    }
                },
                "enforce": {
                    "properties": {
                        "min_validation_period": {
                            "$ref": "http://determined.ai/schemas/expconf/v0/check-epoch-not-used.json"
                        },
                        "min_checkpoint_period": {
                            "$ref": "http://determined.ai/schemas/expconf/v0/check-epoch-not-used.json"
                        },
                        "searcher": {
                            "$ref": "http://determined.ai/schemas/expconf/v0/check-epoch-not-used.json"
                        }
                    }
                }
            }
        }
    ],
    "checks": {
        "must specify an entrypoint that references the trial class": {
            "conditional": {
                "$comment": "when internal.native is null, expect an entrypoint",
                "when": {
                    "properties": {
                        "internal": {
                            "properties": {
                                "native": {
                                    "type": "null"
                                }
                            }
                        }
                    }
                },
                "enforce": {
                    "not": {
                        "properties": {
                            "entrypoint": {
                                "type": "null"
                            }
                        }
                    }
                }
            }
        }
    }
}
`)
	textGCSConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/gcs.json",
    "title": "GCSConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "type",
        "bucket"
    ],
    "properties": {
        "type": {
            "const": "gcs"
        },
        "bucket": {
            "type": "string"
        },
        "save_experiment_best": {
            "type": [
                "integer",
                "null"
            ],
            "default": 0,
            "minimum": 0
        },
        "save_trial_best": {
            "type": [
                "integer",
                "null"
            ],
            "default": 1,
            "minimum": 0
        },
        "save_trial_latest": {
            "type": [
                "integer",
                "null"
            ],
            "default": 1,
            "minimum": 0
        }
    }
}
`)
	textHDFSConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/hdfs.json",
    "title": "HDFSConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "type",
        "hdfs_url",
        "hdfs_path"
    ],
    "properties": {
        "type": {
            "const": "hdfs"
        },
        "hdfs_url": {
            "type": "string"
        },
        "hdfs_path": {
            "type": "string",
            "checks": {
                "hdfs_path must be an absolute path": {
                    "pattern": "^/"
                }
            }
        },
        "user": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        },
        "save_experiment_best": {
            "type": [
                "integer",
                "null"
            ],
            "default": 0,
            "minimum": 0
        },
        "save_trial_best": {
            "type": [
                "integer",
                "null"
            ],
            "default": 1,
            "minimum": 0
        },
        "save_trial_latest": {
            "type": [
                "integer",
                "null"
            ],
            "default": 1,
            "minimum": 0
        }
    }
}
`)
	textCategoricalHyperparameterV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/hyperparameter-categorical.json",
    "title": "CategoricalHyperparameter",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "type",
        "vals"
    ],
    "properties": {
        "type": {
            "const": "categorical"
        },
        "vals": {
            "type": "array",
            "minLength": 1
        }
    }
}
`)
	textConstHyperparameterV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/hyperparameter-const.json",
    "title": "ConstHyperparameter",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "type",
        "val"
    ],
    "properties": {
        "type": {
            "const": "const"
        },
        "val": true
    }
}
`)
	textDoubleHyperparameterV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/hyperparameter-double.json",
    "title": "DoubleHyperparameter",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "type",
        "minval",
        "maxval"
    ],
    "properties": {
        "type": {
            "const": "double"
        },
        "minval": {
            "type": "number"
        },
        "maxval": {
            "type": "number"
        },
        "count": {
            "type": [
                "integer",
                "null"
            ],
            "default": null,
            "minimum": 1
        }
    },
    "compareProperties": {
        "type": "a<b",
        "a": "minval",
        "b": "maxval"
    }
}
`)
	textIntHyperparameterV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/hyperparameter-int.json",
    "title": "IntHyperparameter",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "type",
        "minval",
        "maxval"
    ],
    "properties": {
        "type": {
            "const": "int"
        },
        "minval": {
            "type": "integer"
        },
        "maxval": {
            "type": "integer"
        },
        "count": {
            "type": [
                "integer",
                "null"
            ],
            "default": null,
            "minimum": 1
        }
    },
    "compareProperties": {
        "type": "a<b",
        "a": "minval",
        "b": "maxval"
    }
}
`)
	textLogHyperparameterV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/hyperparameter-log.json",
    "title": "LogHyperparameter",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "type",
        "minval",
        "maxval",
        "base"
    ],
    "properties": {
        "type": {
            "const": "log"
        },
        "minval": {
            "type": "number"
        },
        "maxval": {
            "type": "number"
        },
        "base": {
            "type": "number",
            "exclusiveMinimum": 0
        },
        "count": {
            "type": [
                "integer",
                "null"
            ],
            "default": null,
            "minimum": 1
        }
    },
    "compareProperties": {
        "type": "a<b",
        "a": "minval",
        "b": "maxval"
    }
}
`)
	textHyperparameterV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/hyperparameter.json",
    "title": "Hyperparameter",
    "union": {
        "items": [
            {
                "unionKey": "const:type=int",
                "$ref": "http://determined.ai/schemas/expconf/v0/hyperparameter-int.json"
            },
            {
                "unionKey": "const:type=double",
                "$ref": "http://determined.ai/schemas/expconf/v0/hyperparameter-double.json"
            },
            {
                "unionKey": "const:type=log",
                "$ref": "http://determined.ai/schemas/expconf/v0/hyperparameter-log.json"
            },
            {
                "unionKey": "const:type=const",
                "$ref": "http://determined.ai/schemas/expconf/v0/hyperparameter-const.json"
            },
            {
                "unionKey": "const:type=categorical",
                "$ref": "http://determined.ai/schemas/expconf/v0/hyperparameter-categorical.json"
            },
            {
                "unionKey": "always",
                "type": "object",
                "checks": {
                    "if a hyperparameter object's [\"type\"] is set, it must be one of \"int\", \"double\", \"log\", const\", or \"categorical\"": {
                        "properties": {
                            "type": false
                        }
                    }
                },
                "additionalProperties": {
                    "$ref": "http://determined.ai/schemas/expconf/v0/hyperparameter.json"
                }
            },
            {
                "unionKey": "never",
                "not": {
                    "type": "object"
                }
            }
        ]
    }
}
`)
	textHyperparametersV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/hyperparameters.json",
    "title": "Hyperparameters",
    "type": "object",
    "required": [
        "global_batch_size"
    ],
    "properties": {
        "global_batch_size": {
            "$ref": "http://determined.ai/schemas/expconf/v0/check-global-batch-size.json"
        }
    },
    "additionalProperties": {
        "$ref": "http://determined.ai/schemas/expconf/v0/hyperparameter.json"
    }
}
`)
	textInternalConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/internal.json",
    "title": "InternalConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "native"
    ],
    "properties": {
        "native": {
            "type": "array",
            "items": {
                "type": "string"
            }
        }
    }
}
`)
	textKerberosConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/kerberos.json",
    "title": "KerberosConfig",
    "$comment": "KerberosConfig has not been used in a very long time",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "config_file"
    ],
    "properties": {
        "config_file": {
            "type": "string"
        }
    }
}
`)
	textLengthV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/length.json",
    "title": "Length",
    "union": {
        "defaultMessage": "a length object must have one attribute named \"batches\", \"records\", or \"epochs\"",
        "items": [
            {
                "unionKey": "singleproperty:batches",
                "type": "object",
                "additionalProperties": false,
                "required": [
                    "batches"
                ],
                "properties": {
                    "batches": {
                        "type": "integer",
                        "minimum": 0
                    }
                }
            },
            {
                "unionKey": "singleproperty:records",
                "type": "object",
                "additionalProperties": false,
                "required": [
                    "records"
                ],
                "properties": {
                    "records": {
                        "type": "integer",
                        "minimum": 0
                    }
                }
            },
            {
                "unionKey": "singleproperty:epochs",
                "type": "object",
                "additionalProperties": false,
                "required": [
                    "epochs"
                ],
                "properties": {
                    "epochs": {
                        "type": "integer",
                        "minimum": 0
                    }
                }
            }
        ]
    }
}
`)
	textOptimizationsConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/optimizations.json",
    "title": "OptimizationsConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [],
    "properties": {
        "aggregation_frequency": {
            "type": [
                "integer",
                "null"
            ],
            "minimum": 1,
            "default": 1
        },
        "auto_tune_tensor_fusion": {
            "type": [
                "boolean",
                "null"
            ],
            "default": false
        },
        "average_aggregated_gradients": {
            "type": [
                "boolean",
                "null"
            ],
            "default": true
        },
        "average_training_metrics": {
            "type": [
                "boolean",
                "null"
            ],
            "default": false
        },
        "gradient_compression": {
            "type": [
                "boolean",
                "null"
            ],
            "default": false
        },
        "mixed_precision": {
            "enum": [
                null,
                "O0",
                "O1",
                "O2",
                "O3"
            ],
            "default": "O0",
            "checks": {
                "mixed_precision should be a string starting with an uppercase letter 'O'": {
                    "pattern": "^O"
                }
            }
        },
        "tensor_fusion_cycle_time": {
            "type": [
                "integer",
                "null"
            ],
            "minimum": 0,
            "default": 5
        },
        "tensor_fusion_threshold": {
            "type": [
                "integer",
                "null"
            ],
            "minimum": 0,
            "default": 64
        }
    }
}
`)
	textReproducibilityConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/reproducibility.json",
    "title": "ReproducibilityConfig",
    "type": "object",
    "additionalProperties": false,
    "properties": {
        "experiment_seed": {
            "type": [
                "integer",
                "null"
            ],
            "default": null
        }
    }
}
`)
	textResourcesConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/resources.json",
    "title": "ResourcesConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [],
    "properties": {
        "agent_label": {
            "type": [
                "string",
                "null"
            ],
            "default": ""
        },
        "max_slots": {
            "type": [
                "integer",
                "null"
            ],
            "default": null
        },
        "native_parallel": {
            "type": [
                "boolean",
                "null"
            ],
            "default": false
        },
        "priority": {
            "type": [
                "integer",
                "null"
            ],
            "minimum": 1,
            "maximum": 99,
            "default": null
        },
        "resource_pool": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        },
        "shm_size": {
            "type": [
                "integer",
                "null"
            ],
            "default": null
        },
        "slots_per_trial": {
            "type": [
                "integer",
                "null"
            ],
            "minimum": 0,
            "default": 1
        },
        "weight": {
            "type": [
                "number",
                "null"
            ],
            "default": 1
        }
    }
}
`)
	textS3ConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/s3.json",
    "title": "S3Config",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "type",
        "access_key",
        "bucket",
        "secret_key"
    ],
    "properties": {
        "type": {
            "const": "s3"
        },
        "access_key": {
            "type": "string"
        },
        "bucket": {
            "type": "string"
        },
        "secret_key": {
            "type": "string"
        },
        "endpoint_url": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        },
        "save_experiment_best": {
            "type": [
                "integer",
                "null"
            ],
            "default": 0,
            "minimum": 0
        },
        "save_trial_best": {
            "type": [
                "integer",
                "null"
            ],
            "default": 1,
            "minimum": 0
        },
        "save_trial_latest": {
            "type": [
                "integer",
                "null"
            ],
            "default": 1,
            "minimum": 0
        }
    }
}
`)
	textAdaptiveASHAConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/searcher-adaptive-asha.json",
    "title": "AdaptiveASHAConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "name",
        "max_length",
        "max_trials",
        "metric"
    ],
    "properties": {
        "name": {
            "const": "adaptive_asha"
        },
        "bracket_rungs": {
            "type": [
                "array",
                "null"
            ],
            "default": [],
            "items": {
                "type": "integer"
            }
        },
        "max_trials": {
            "type": "integer",
            "minimum": 1
        },
        "mode": {
            "enum": [
                null,
                "aggressive",
                "standard",
                "conservative"
            ],
            "default": "standard"
        },
        "divisor": {
            "type": [
                "number",
                "null"
            ],
            "exclusiveMinimum": 1,
            "default": 4
        },
        "max_rungs": {
            "type": [
                "integer",
                "null"
            ],
            "minimum": 1,
            "default": 5
        },
        "max_concurrent_trials": {
            "type": [
                "integer",
                "null"
            ],
            "minimum": 0,
            "default": 0
        },
        "max_length": {
            "$ref": "http://determined.ai/schemas/expconf/v0/check-positive-length.json"
        },
        "metric": {
            "type": "string"
        },
        "smaller_is_better": {
            "type": [
                "boolean",
                "null"
            ],
            "default": true
        },
        "source_trial_id": {
            "type": [
                "integer",
                "null"
            ],
            "default": null
        },
        "source_checkpoint_uuid": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        }
    }
}
`)
	textAdaptiveSimpleConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/searcher-adaptive-simple.json",
    "title": "AdaptiveSimpleConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "name",
        "max_trials",
        "max_length",
        "metric"
    ],
    "properties": {
        "name": {
            "const": "adaptive_simple"
        },
        "max_trials": {
            "type": "integer",
            "minimum": 1,
            "maximum": 2000
        },
        "mode": {
            "enum": [
                null,
                "aggressive",
                "standard",
                "conservative"
            ],
            "default": "standard"
        },
        "divisor": {
            "type": [
                "number",
                "null"
            ],
            "exclusiveMinimum": 1,
            "default": 4
        },
        "max_rungs": {
            "type": [
                "integer",
                "null"
            ],
            "minimum": 1,
            "default": 5
        },
        "max_length": {
            "$ref": "http://determined.ai/schemas/expconf/v0/check-positive-length.json"
        },
        "metric": {
            "type": "string"
        },
        "smaller_is_better": {
            "type": [
                "boolean",
                "null"
            ],
            "default": true
        },
        "source_trial_id": {
            "type": [
                "integer",
                "null"
            ],
            "default": null
        },
        "source_checkpoint_uuid": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        }
    }
}
`)
	textAdaptiveConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/searcher-adaptive.json",
    "title": "AdaptiveConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "name",
        "budget",
        "max_length",
        "metric"
    ],
    "properties": {
        "name": {
            "const": "adaptive"
        },
        "budget": {
            "$ref": "http://determined.ai/schemas/expconf/v0/length.json"
        },
        "bracket_rungs": {
            "type": [
                "array",
                "null"
            ],
            "default": [],
            "items": {
                "type": "integer"
            }
        },
        "mode": {
            "enum": [
                null,
                "aggressive",
                "standard",
                "conservative"
            ],
            "default": "standard"
        },
        "divisor": {
            "type": [
                "number",
                "null"
            ],
            "exclusiveMinimum": 1,
            "default": 4
        },
        "max_rungs": {
            "type": [
                "integer",
                "null"
            ],
            "minimum": 1,
            "default": 5
        },
        "max_length": {
            "$ref": "http://determined.ai/schemas/expconf/v0/check-positive-length.json"
        },
        "train_stragglers": {
            "type": [
                "boolean",
                "null"
            ],
            "default": true
        },
        "metric": {
            "type": "string"
        },
        "smaller_is_better": {
            "type": [
                "boolean",
                "null"
            ],
            "default": true
        },
        "source_trial_id": {
            "type": [
                "integer",
                "null"
            ],
            "default": null
        },
        "source_checkpoint_uuid": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        }
    },
    "checks": {
        "max_length and budget must be specified in terms of the same unit": {
            "compareProperties": {
                "type": "same_units",
                "a": "max_length",
                "b": "budget"
            }
        },
        "budget must be greater than max_length": {
            "compareProperties": {
                "type": "length_a<length_b",
                "a": "max_length",
                "b": "budget"
            }
        }
    }
}
`)
	textAsyncHalvingConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/searcher-async-halving.json",
    "title": "AsyncHalvingConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "name",
        "num_rungs",
        "max_length",
        "max_trials",
        "metric"
    ],
    "properties": {
        "name": {
            "const": "async_halving"
        },
        "num_rungs": {
            "type": "integer",
            "minimum": 1
        },
        "max_length": {
            "$ref": "http://determined.ai/schemas/expconf/v0/check-positive-length.json"
        },
        "max_trials": {
            "type": "integer",
            "minimum": 1
        },
        "divisor": {
            "type": [
                "number",
                "null"
            ],
            "exclusiveMinimum": 1,
            "default": 4
        },
        "max_concurrent_trials": {
            "type": [
                "integer",
                "null"
            ],
            "minimum": 0,
            "default": 0
        },
        "metric": {
            "type": "string"
        },
        "smaller_is_better": {
            "type": [
                "boolean",
                "null"
            ],
            "default": true
        },
        "source_trial_id": {
            "type": [
                "integer",
                "null"
            ],
            "default": null
        },
        "source_checkpoint_uuid": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        }
    }
}
`)
	textGridConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/searcher-grid.json",
    "title": "GridConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "name",
        "max_length",
        "metric"
    ],
    "properties": {
        "name": {
            "const": "grid"
        },
        "max_length": {
            "$ref": "http://determined.ai/schemas/expconf/v0/check-positive-length.json"
        },
        "metric": {
            "type": "string"
        },
        "smaller_is_better": {
            "type": [
                "boolean",
                "null"
            ],
            "default": true
        },
        "source_trial_id": {
            "type": [
                "integer",
                "null"
            ],
            "default": null
        },
        "source_checkpoint_uuid": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        }
    }
}
`)
	textPBTConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/searcher-pbt.json",
    "title": "PBTConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "name",
        "metric",
        "population_size",
        "length_per_round",
        "num_rounds",
        "replace_function",
        "explore_function"
    ],
    "properties": {
        "name": {
            "const": "pbt"
        },
        "population_size": {
            "type": "integer",
            "minimum": 1
        },
        "length_per_round": {
            "$ref": "http://determined.ai/schemas/expconf/v0/check-positive-length.json"
        },
        "num_rounds": {
            "type": "integer",
            "minimum": 1
        },
        "replace_function": {
            "unionKey": "singleproperty",
            "union": {
                "items": [
                    {
                        "unionKey": "always",
                        "type": "object",
                        "additionalProperties": false,
                        "required": [
                            "truncate_fraction"
                        ],
                        "properties": {
                            "truncate_fraction": {
                                "type": "number",
                                "minimum": 0.0,
                                "maximum": 0.5
                            }
                        }
                    }
                ]
            }
        },
        "explore_function": {
            "type": "object",
            "additionalProperties": false,
            "required": [
                "resample_probability",
                "perturb_factor"
            ],
            "properties": {
                "resample_probability": {
                    "type": "number",
                    "minimum": 0.0,
                    "maximum": 1.0
                },
                "perturb_factor": {
                    "type": "number",
                    "minimum": 0.0,
                    "maximum": 1.0
                }
            }
        },
        "metric": {
            "type": "string"
        },
        "smaller_is_better": {
            "type": [
                "boolean",
                "null"
            ],
            "default": true
        },
        "source_trial_id": {
            "type": [
                "integer",
                "null"
            ],
            "default": null
        },
        "source_checkpoint_uuid": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        }
    }
}
`)
	textRandomConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/searcher-random.json",
    "title": "RandomConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "name",
        "max_trials",
        "max_length",
        "metric"
    ],
    "properties": {
        "name": {
            "const": "random"
        },
        "max_trials": {
            "type": "integer",
            "minimum": 1
        },
        "max_length": {
            "$ref": "http://determined.ai/schemas/expconf/v0/check-positive-length.json"
        },
        "metric": {
            "type": "string"
        },
        "smaller_is_better": {
            "type": [
                "boolean",
                "null"
            ],
            "default": true
        },
        "source_trial_id": {
            "type": [
                "integer",
                "null"
            ],
            "default": null
        },
        "source_checkpoint_uuid": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        }
    }
}
`)
	textSingleConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/searcher-single.json",
    "title": "SingleConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "name",
        "max_length",
        "metric"
    ],
    "properties": {
        "name": {
            "const": "single"
        },
        "max_length": {
            "$ref": "http://determined.ai/schemas/expconf/v0/check-positive-length.json"
        },
        "metric": {
            "type": "string"
        },
        "smaller_is_better": {
            "type": [
                "boolean",
                "null"
            ],
            "default": true
        },
        "source_trial_id": {
            "type": [
                "integer",
                "null"
            ],
            "default": null
        },
        "source_checkpoint_uuid": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        }
    }
}
`)
	textSyncHalvingConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/searcher-sync-halving.json",
    "title": "SyncHalvingConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "name",
        "num_rungs",
        "max_length",
        "budget",
        "metric"
    ],
    "properties": {
        "name": {
            "const": "sync_halving"
        },
        "budget": {
            "$ref": "http://determined.ai/schemas/expconf/v0/check-positive-length.json"
        },
        "num_rungs": {
            "type": "integer",
            "minimum": 1
        },
        "max_length": {
            "$ref": "http://determined.ai/schemas/expconf/v0/check-positive-length.json"
        },
        "divisor": {
            "type": [
                "number",
                "null"
            ],
            "exclusiveMinimum": 1,
            "default": 4
        },
        "train_stragglers": {
            "type": [
                "boolean",
                "null"
            ],
            "default": true
        },
        "metric": {
            "type": "string"
        },
        "smaller_is_better": {
            "type": [
                "boolean",
                "null"
            ],
            "default": true
        },
        "source_trial_id": {
            "type": [
                "integer",
                "null"
            ],
            "default": null
        },
        "source_checkpoint_uuid": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        }
    }
}
`)
	textSearcherConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/searcher.json",
    "title": "SearcherConfig",
    "union": {
        "defaultMessage": "is not an object where object[\"name\"] is one of 'single', 'random', 'grid', 'adaptive', 'adaptive_asha', 'adaptive_simple', or 'pbt'",
        "items": [
            {
                "unionKey": "const:name=single",
                "$ref": "http://determined.ai/schemas/expconf/v0/searcher-single.json"
            },
            {
                "unionKey": "const:name=random",
                "$ref": "http://determined.ai/schemas/expconf/v0/searcher-random.json"
            },
            {
                "unionKey": "const:name=grid",
                "$ref": "http://determined.ai/schemas/expconf/v0/searcher-grid.json"
            },
            {
                "unionKey": "const:name=adaptive_asha",
                "$ref": "http://determined.ai/schemas/expconf/v0/searcher-adaptive-asha.json"
            },
            {
                "unionKey": "const:name=adaptive_simple",
                "$ref": "http://determined.ai/schemas/expconf/v0/searcher-adaptive-simple.json"
            },
            {
                "unionKey": "const:name=adaptive",
                "$ref": "http://determined.ai/schemas/expconf/v0/searcher-adaptive.json"
            },
            {
                "unionKey": "const:name=pbt",
                "$ref": "http://determined.ai/schemas/expconf/v0/searcher-pbt.json"
            },
            {
                "unionKey": "const:name=sync_halving",
                "$ref": "http://determined.ai/schemas/expconf/v0/searcher-sync-halving.json"
            },
            {
                "unionKey": "const:name=async_halving",
                "$ref": "http://determined.ai/schemas/expconf/v0/searcher-async-halving.json"
            }
        ]
    }
}
`)
	textSecurityConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/security.json",
    "title": "SecurityConfig",
    "$comment": "SecurityConfig has not been used in a very long time",
    "type": "object",
    "additionalProperties": false,
    "properties": {
        "kerberos": {
            "type": [
                "object",
                "null"
            ],
            "default": null,
            "optionalRef": "http://determined.ai/schemas/expconf/v0/kerberos.json"
        }
    }
}
`)
	textSharedFSConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/shared-fs.json",
    "title": "SharedFSConfig",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "type",
        "host_path"
    ],
    "properties": {
        "type": {
            "const": "shared_fs"
        },
        "host_path": {
            "type": "string"
        },
        "storage_path": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        },
        "propagation": {
            "type": [
                "string",
                "null"
            ],
            "default": "rprivate"
        },
        "container_path": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        },
        "checkpoint_path": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        },
        "tensorboard_path": {
            "type": [
                "string",
                "null"
            ],
            "default": null
        },
        "save_experiment_best": {
            "type": [
                "integer",
                "null"
            ],
            "default": 0,
            "minimum": 0
        },
        "save_trial_best": {
            "type": [
                "integer",
                "null"
            ],
            "default": 1,
            "minimum": 0
        },
        "save_trial_latest": {
            "type": [
                "integer",
                "null"
            ],
            "default": 1,
            "minimum": 0
        }
    },
    "checks": {
        "storage_path must either be a relative directory or a subdirectory of host_path": {
            "compareProperties": {
                "type": "a_is_subdir_of_b",
                "a": "storage_path",
                "b": "host_path"
            }
        }
    }
}
`)
	textTensorboardStorageConfigV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/tensorboard-storage.json",
    "title": "TensorboardStorageConfig",
    "$comment": "TensorboardStorageConfig has not been used in a very long time",
    "union": {
        "defaultMessage": "this field is deprecated and will be ignored",
        "items": [
            {
                "unionKey": "never",
                "$ref": "http://determined.ai/schemas/expconf/v0/shared-fs.json"
            },
            {
                "unionKey": "never",
                "$ref": "http://determined.ai/schemas/expconf/v0/hdfs.json"
            },
            {
                "unionKey": "never",
                "$ref": "http://determined.ai/schemas/expconf/v0/s3.json"
            },
            {
                "unionKey": "never",
                "$ref": "http://determined.ai/schemas/expconf/v0/gcs.json"
            }
        ]
    },
    "disallowProperties": {
        "save_experiment_best": "this field is deprecated and will be ignored",
        "save_trial_best": "this field is deprecated and will be ignored",
        "save_trial_latest": "this field is deprecated and will be ignored"
    }
}
`)
	textTestRootV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/test-root.json",
    "title": "TestRoot",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "val_x"
    ],
    "properties": {
        "val_x": {
            "type": "integer"
        },
        "sub_obj": {
            "type": [
                "object",
                "null"
            ],
            "default": {},
            "optionalRef": "http://determined.ai/schemas/expconf/v0/test-sub.json"
        },
        "sub_union": {
            "type": [
                "object",
                "null"
            ],
            "default": null,
            "optionalRef": "http://determined.ai/schemas/expconf/v0/test-union.json"
        },
        "runtime_defaultable": {
            "type": [
                "integer",
                "null"
            ],
            "default": null
        }
    }
}
`)
	textTestSubV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/test-sub.json",
    "title": "TestSub",
    "type": "object",
    "additionalProperties": false,
    "required": [],
    "properties": {
        "val_y": {
            "type": [
                "string",
                "null"
            ],
            "default": "default_y"
        }
    }
}
`)
	textTestUnionAV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/test-union-a.json",
    "title": "TestUnionA",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "type",
        "val_a"
    ],
    "properties": {
        "type": {
            "const": "a"
        },
        "val_a": {
            "type": "integer"
        },
        "common_val": {
            "type": [
                "string",
                "null"
            ],
            "default": "default-common-val"
        }
    }
}
`)
	textTestUnionBV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/test-union-b.json",
    "title": "TestUnionB",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "type",
        "val_b"
    ],
    "properties": {
        "type": {
            "const": "b"
        },
        "val_b": {
            "type": "integer"
        },
        "common_val": {
            "type": [
                "string",
                "null"
            ],
            "default": "default-common-val"
        }
    }
}
`)
	textTestUnionV0 = []byte(`{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "$id": "http://determined.ai/schemas/expconf/v0/test-union.json",
    "title": "TestUnion",
    "union": {
        "defaultMessage": "bad test union",
        "items": [
            {
                "unionKey": "const:type=a",
                "$ref": "http://determined.ai/schemas/expconf/v0/test-union-a.json"
            },
            {
                "unionKey": "const:type=b",
                "$ref": "http://determined.ai/schemas/expconf/v0/test-union-b.json"
            }
        ]
    }
}
`)
	schemaBindMountV0                 interface{}
	schemaCheckDataLayerCacheV0       interface{}
	schemaCheckEpochNotUsedV0         interface{}
	schemaCheckGlobalBatchSizeV0      interface{}
	schemaCheckGridHyperparameterV0   interface{}
	schemaCheckPositiveLengthV0       interface{}
	schemaCheckpointStorageConfigV0   interface{}
	schemaGCSDataLayerConfigV0        interface{}
	schemaS3DataLayerConfigV0         interface{}
	schemaSharedFSDataLayerConfigV0   interface{}
	schemaDataLayerConfigV0           interface{}
	schemaEnvironmentImageMapV0       interface{}
	schemaEnvironmentImageV0          interface{}
	schemaEnvironmentVariablesMapV0   interface{}
	schemaEnvironmentVariablesV0      interface{}
	schemaEnvironmentConfigV0         interface{}
	schemaExperimentConfigV0          interface{}
	schemaGCSConfigV0                 interface{}
	schemaHDFSConfigV0                interface{}
	schemaCategoricalHyperparameterV0 interface{}
	schemaConstHyperparameterV0       interface{}
	schemaDoubleHyperparameterV0      interface{}
	schemaIntHyperparameterV0         interface{}
	schemaLogHyperparameterV0         interface{}
	schemaHyperparameterV0            interface{}
	schemaHyperparametersV0           interface{}
	schemaInternalConfigV0            interface{}
	schemaKerberosConfigV0            interface{}
	schemaLengthV0                    interface{}
	schemaOptimizationsConfigV0       interface{}
	schemaReproducibilityConfigV0     interface{}
	schemaResourcesConfigV0           interface{}
	schemaS3ConfigV0                  interface{}
	schemaAdaptiveASHAConfigV0        interface{}
	schemaAdaptiveSimpleConfigV0      interface{}
	schemaAdaptiveConfigV0            interface{}
	schemaAsyncHalvingConfigV0        interface{}
	schemaGridConfigV0                interface{}
	schemaPBTConfigV0                 interface{}
	schemaRandomConfigV0              interface{}
	schemaSingleConfigV0              interface{}
	schemaSyncHalvingConfigV0         interface{}
	schemaSearcherConfigV0            interface{}
	schemaSecurityConfigV0            interface{}
	schemaSharedFSConfigV0            interface{}
	schemaTensorboardStorageConfigV0  interface{}
	schemaTestRootV0                  interface{}
	schemaTestSubV0                   interface{}
	schemaTestUnionAV0                interface{}
	schemaTestUnionBV0                interface{}
	schemaTestUnionV0                 interface{}
	cachedSchemaMap                   map[string]interface{}
	cachedSchemaBytesMap              map[string][]byte
)

func ParsedBindMountV0() interface{} {
	if schemaBindMountV0 != nil {
		return schemaBindMountV0
	}
	err := json.Unmarshal(textBindMountV0, &schemaBindMountV0)
	if err != nil {
		panic("invalid embedded json for BindMountV0")
	}
	return schemaBindMountV0
}

func ParsedCheckDataLayerCacheV0() interface{} {
	if schemaCheckDataLayerCacheV0 != nil {
		return schemaCheckDataLayerCacheV0
	}
	err := json.Unmarshal(textCheckDataLayerCacheV0, &schemaCheckDataLayerCacheV0)
	if err != nil {
		panic("invalid embedded json for CheckDataLayerCacheV0")
	}
	return schemaCheckDataLayerCacheV0
}

func ParsedCheckEpochNotUsedV0() interface{} {
	if schemaCheckEpochNotUsedV0 != nil {
		return schemaCheckEpochNotUsedV0
	}
	err := json.Unmarshal(textCheckEpochNotUsedV0, &schemaCheckEpochNotUsedV0)
	if err != nil {
		panic("invalid embedded json for CheckEpochNotUsedV0")
	}
	return schemaCheckEpochNotUsedV0
}

func ParsedCheckGlobalBatchSizeV0() interface{} {
	if schemaCheckGlobalBatchSizeV0 != nil {
		return schemaCheckGlobalBatchSizeV0
	}
	err := json.Unmarshal(textCheckGlobalBatchSizeV0, &schemaCheckGlobalBatchSizeV0)
	if err != nil {
		panic("invalid embedded json for CheckGlobalBatchSizeV0")
	}
	return schemaCheckGlobalBatchSizeV0
}

func ParsedCheckGridHyperparameterV0() interface{} {
	if schemaCheckGridHyperparameterV0 != nil {
		return schemaCheckGridHyperparameterV0
	}
	err := json.Unmarshal(textCheckGridHyperparameterV0, &schemaCheckGridHyperparameterV0)
	if err != nil {
		panic("invalid embedded json for CheckGridHyperparameterV0")
	}
	return schemaCheckGridHyperparameterV0
}

func ParsedCheckPositiveLengthV0() interface{} {
	if schemaCheckPositiveLengthV0 != nil {
		return schemaCheckPositiveLengthV0
	}
	err := json.Unmarshal(textCheckPositiveLengthV0, &schemaCheckPositiveLengthV0)
	if err != nil {
		panic("invalid embedded json for CheckPositiveLengthV0")
	}
	return schemaCheckPositiveLengthV0
}

func ParsedCheckpointStorageConfigV0() interface{} {
	if schemaCheckpointStorageConfigV0 != nil {
		return schemaCheckpointStorageConfigV0
	}
	err := json.Unmarshal(textCheckpointStorageConfigV0, &schemaCheckpointStorageConfigV0)
	if err != nil {
		panic("invalid embedded json for CheckpointStorageConfigV0")
	}
	return schemaCheckpointStorageConfigV0
}

func ParsedGCSDataLayerConfigV0() interface{} {
	if schemaGCSDataLayerConfigV0 != nil {
		return schemaGCSDataLayerConfigV0
	}
	err := json.Unmarshal(textGCSDataLayerConfigV0, &schemaGCSDataLayerConfigV0)
	if err != nil {
		panic("invalid embedded json for GCSDataLayerConfigV0")
	}
	return schemaGCSDataLayerConfigV0
}

func ParsedS3DataLayerConfigV0() interface{} {
	if schemaS3DataLayerConfigV0 != nil {
		return schemaS3DataLayerConfigV0
	}
	err := json.Unmarshal(textS3DataLayerConfigV0, &schemaS3DataLayerConfigV0)
	if err != nil {
		panic("invalid embedded json for S3DataLayerConfigV0")
	}
	return schemaS3DataLayerConfigV0
}

func ParsedSharedFSDataLayerConfigV0() interface{} {
	if schemaSharedFSDataLayerConfigV0 != nil {
		return schemaSharedFSDataLayerConfigV0
	}
	err := json.Unmarshal(textSharedFSDataLayerConfigV0, &schemaSharedFSDataLayerConfigV0)
	if err != nil {
		panic("invalid embedded json for SharedFSDataLayerConfigV0")
	}
	return schemaSharedFSDataLayerConfigV0
}

func ParsedDataLayerConfigV0() interface{} {
	if schemaDataLayerConfigV0 != nil {
		return schemaDataLayerConfigV0
	}
	err := json.Unmarshal(textDataLayerConfigV0, &schemaDataLayerConfigV0)
	if err != nil {
		panic("invalid embedded json for DataLayerConfigV0")
	}
	return schemaDataLayerConfigV0
}

func ParsedEnvironmentImageMapV0() interface{} {
	if schemaEnvironmentImageMapV0 != nil {
		return schemaEnvironmentImageMapV0
	}
	err := json.Unmarshal(textEnvironmentImageMapV0, &schemaEnvironmentImageMapV0)
	if err != nil {
		panic("invalid embedded json for EnvironmentImageMapV0")
	}
	return schemaEnvironmentImageMapV0
}

func ParsedEnvironmentImageV0() interface{} {
	if schemaEnvironmentImageV0 != nil {
		return schemaEnvironmentImageV0
	}
	err := json.Unmarshal(textEnvironmentImageV0, &schemaEnvironmentImageV0)
	if err != nil {
		panic("invalid embedded json for EnvironmentImageV0")
	}
	return schemaEnvironmentImageV0
}

func ParsedEnvironmentVariablesMapV0() interface{} {
	if schemaEnvironmentVariablesMapV0 != nil {
		return schemaEnvironmentVariablesMapV0
	}
	err := json.Unmarshal(textEnvironmentVariablesMapV0, &schemaEnvironmentVariablesMapV0)
	if err != nil {
		panic("invalid embedded json for EnvironmentVariablesMapV0")
	}
	return schemaEnvironmentVariablesMapV0
}

func ParsedEnvironmentVariablesV0() interface{} {
	if schemaEnvironmentVariablesV0 != nil {
		return schemaEnvironmentVariablesV0
	}
	err := json.Unmarshal(textEnvironmentVariablesV0, &schemaEnvironmentVariablesV0)
	if err != nil {
		panic("invalid embedded json for EnvironmentVariablesV0")
	}
	return schemaEnvironmentVariablesV0
}

func ParsedEnvironmentConfigV0() interface{} {
	if schemaEnvironmentConfigV0 != nil {
		return schemaEnvironmentConfigV0
	}
	err := json.Unmarshal(textEnvironmentConfigV0, &schemaEnvironmentConfigV0)
	if err != nil {
		panic("invalid embedded json for EnvironmentConfigV0")
	}
	return schemaEnvironmentConfigV0
}

func ParsedExperimentConfigV0() interface{} {
	if schemaExperimentConfigV0 != nil {
		return schemaExperimentConfigV0
	}
	err := json.Unmarshal(textExperimentConfigV0, &schemaExperimentConfigV0)
	if err != nil {
		panic("invalid embedded json for ExperimentConfigV0")
	}
	return schemaExperimentConfigV0
}

func ParsedGCSConfigV0() interface{} {
	if schemaGCSConfigV0 != nil {
		return schemaGCSConfigV0
	}
	err := json.Unmarshal(textGCSConfigV0, &schemaGCSConfigV0)
	if err != nil {
		panic("invalid embedded json for GCSConfigV0")
	}
	return schemaGCSConfigV0
}

func ParsedHDFSConfigV0() interface{} {
	if schemaHDFSConfigV0 != nil {
		return schemaHDFSConfigV0
	}
	err := json.Unmarshal(textHDFSConfigV0, &schemaHDFSConfigV0)
	if err != nil {
		panic("invalid embedded json for HDFSConfigV0")
	}
	return schemaHDFSConfigV0
}

func ParsedCategoricalHyperparameterV0() interface{} {
	if schemaCategoricalHyperparameterV0 != nil {
		return schemaCategoricalHyperparameterV0
	}
	err := json.Unmarshal(textCategoricalHyperparameterV0, &schemaCategoricalHyperparameterV0)
	if err != nil {
		panic("invalid embedded json for CategoricalHyperparameterV0")
	}
	return schemaCategoricalHyperparameterV0
}

func ParsedConstHyperparameterV0() interface{} {
	if schemaConstHyperparameterV0 != nil {
		return schemaConstHyperparameterV0
	}
	err := json.Unmarshal(textConstHyperparameterV0, &schemaConstHyperparameterV0)
	if err != nil {
		panic("invalid embedded json for ConstHyperparameterV0")
	}
	return schemaConstHyperparameterV0
}

func ParsedDoubleHyperparameterV0() interface{} {
	if schemaDoubleHyperparameterV0 != nil {
		return schemaDoubleHyperparameterV0
	}
	err := json.Unmarshal(textDoubleHyperparameterV0, &schemaDoubleHyperparameterV0)
	if err != nil {
		panic("invalid embedded json for DoubleHyperparameterV0")
	}
	return schemaDoubleHyperparameterV0
}

func ParsedIntHyperparameterV0() interface{} {
	if schemaIntHyperparameterV0 != nil {
		return schemaIntHyperparameterV0
	}
	err := json.Unmarshal(textIntHyperparameterV0, &schemaIntHyperparameterV0)
	if err != nil {
		panic("invalid embedded json for IntHyperparameterV0")
	}
	return schemaIntHyperparameterV0
}

func ParsedLogHyperparameterV0() interface{} {
	if schemaLogHyperparameterV0 != nil {
		return schemaLogHyperparameterV0
	}
	err := json.Unmarshal(textLogHyperparameterV0, &schemaLogHyperparameterV0)
	if err != nil {
		panic("invalid embedded json for LogHyperparameterV0")
	}
	return schemaLogHyperparameterV0
}

func ParsedHyperparameterV0() interface{} {
	if schemaHyperparameterV0 != nil {
		return schemaHyperparameterV0
	}
	err := json.Unmarshal(textHyperparameterV0, &schemaHyperparameterV0)
	if err != nil {
		panic("invalid embedded json for HyperparameterV0")
	}
	return schemaHyperparameterV0
}

func ParsedHyperparametersV0() interface{} {
	if schemaHyperparametersV0 != nil {
		return schemaHyperparametersV0
	}
	err := json.Unmarshal(textHyperparametersV0, &schemaHyperparametersV0)
	if err != nil {
		panic("invalid embedded json for HyperparametersV0")
	}
	return schemaHyperparametersV0
}

func ParsedInternalConfigV0() interface{} {
	if schemaInternalConfigV0 != nil {
		return schemaInternalConfigV0
	}
	err := json.Unmarshal(textInternalConfigV0, &schemaInternalConfigV0)
	if err != nil {
		panic("invalid embedded json for InternalConfigV0")
	}
	return schemaInternalConfigV0
}

func ParsedKerberosConfigV0() interface{} {
	if schemaKerberosConfigV0 != nil {
		return schemaKerberosConfigV0
	}
	err := json.Unmarshal(textKerberosConfigV0, &schemaKerberosConfigV0)
	if err != nil {
		panic("invalid embedded json for KerberosConfigV0")
	}
	return schemaKerberosConfigV0
}

func ParsedLengthV0() interface{} {
	if schemaLengthV0 != nil {
		return schemaLengthV0
	}
	err := json.Unmarshal(textLengthV0, &schemaLengthV0)
	if err != nil {
		panic("invalid embedded json for LengthV0")
	}
	return schemaLengthV0
}

func ParsedOptimizationsConfigV0() interface{} {
	if schemaOptimizationsConfigV0 != nil {
		return schemaOptimizationsConfigV0
	}
	err := json.Unmarshal(textOptimizationsConfigV0, &schemaOptimizationsConfigV0)
	if err != nil {
		panic("invalid embedded json for OptimizationsConfigV0")
	}
	return schemaOptimizationsConfigV0
}

func ParsedReproducibilityConfigV0() interface{} {
	if schemaReproducibilityConfigV0 != nil {
		return schemaReproducibilityConfigV0
	}
	err := json.Unmarshal(textReproducibilityConfigV0, &schemaReproducibilityConfigV0)
	if err != nil {
		panic("invalid embedded json for ReproducibilityConfigV0")
	}
	return schemaReproducibilityConfigV0
}

func ParsedResourcesConfigV0() interface{} {
	if schemaResourcesConfigV0 != nil {
		return schemaResourcesConfigV0
	}
	err := json.Unmarshal(textResourcesConfigV0, &schemaResourcesConfigV0)
	if err != nil {
		panic("invalid embedded json for ResourcesConfigV0")
	}
	return schemaResourcesConfigV0
}

func ParsedS3ConfigV0() interface{} {
	if schemaS3ConfigV0 != nil {
		return schemaS3ConfigV0
	}
	err := json.Unmarshal(textS3ConfigV0, &schemaS3ConfigV0)
	if err != nil {
		panic("invalid embedded json for S3ConfigV0")
	}
	return schemaS3ConfigV0
}

func ParsedAdaptiveASHAConfigV0() interface{} {
	if schemaAdaptiveASHAConfigV0 != nil {
		return schemaAdaptiveASHAConfigV0
	}
	err := json.Unmarshal(textAdaptiveASHAConfigV0, &schemaAdaptiveASHAConfigV0)
	if err != nil {
		panic("invalid embedded json for AdaptiveASHAConfigV0")
	}
	return schemaAdaptiveASHAConfigV0
}

func ParsedAdaptiveSimpleConfigV0() interface{} {
	if schemaAdaptiveSimpleConfigV0 != nil {
		return schemaAdaptiveSimpleConfigV0
	}
	err := json.Unmarshal(textAdaptiveSimpleConfigV0, &schemaAdaptiveSimpleConfigV0)
	if err != nil {
		panic("invalid embedded json for AdaptiveSimpleConfigV0")
	}
	return schemaAdaptiveSimpleConfigV0
}

func ParsedAdaptiveConfigV0() interface{} {
	if schemaAdaptiveConfigV0 != nil {
		return schemaAdaptiveConfigV0
	}
	err := json.Unmarshal(textAdaptiveConfigV0, &schemaAdaptiveConfigV0)
	if err != nil {
		panic("invalid embedded json for AdaptiveConfigV0")
	}
	return schemaAdaptiveConfigV0
}

func ParsedAsyncHalvingConfigV0() interface{} {
	if schemaAsyncHalvingConfigV0 != nil {
		return schemaAsyncHalvingConfigV0
	}
	err := json.Unmarshal(textAsyncHalvingConfigV0, &schemaAsyncHalvingConfigV0)
	if err != nil {
		panic("invalid embedded json for AsyncHalvingConfigV0")
	}
	return schemaAsyncHalvingConfigV0
}

func ParsedGridConfigV0() interface{} {
	if schemaGridConfigV0 != nil {
		return schemaGridConfigV0
	}
	err := json.Unmarshal(textGridConfigV0, &schemaGridConfigV0)
	if err != nil {
		panic("invalid embedded json for GridConfigV0")
	}
	return schemaGridConfigV0
}

func ParsedPBTConfigV0() interface{} {
	if schemaPBTConfigV0 != nil {
		return schemaPBTConfigV0
	}
	err := json.Unmarshal(textPBTConfigV0, &schemaPBTConfigV0)
	if err != nil {
		panic("invalid embedded json for PBTConfigV0")
	}
	return schemaPBTConfigV0
}

func ParsedRandomConfigV0() interface{} {
	if schemaRandomConfigV0 != nil {
		return schemaRandomConfigV0
	}
	err := json.Unmarshal(textRandomConfigV0, &schemaRandomConfigV0)
	if err != nil {
		panic("invalid embedded json for RandomConfigV0")
	}
	return schemaRandomConfigV0
}

func ParsedSingleConfigV0() interface{} {
	if schemaSingleConfigV0 != nil {
		return schemaSingleConfigV0
	}
	err := json.Unmarshal(textSingleConfigV0, &schemaSingleConfigV0)
	if err != nil {
		panic("invalid embedded json for SingleConfigV0")
	}
	return schemaSingleConfigV0
}

func ParsedSyncHalvingConfigV0() interface{} {
	if schemaSyncHalvingConfigV0 != nil {
		return schemaSyncHalvingConfigV0
	}
	err := json.Unmarshal(textSyncHalvingConfigV0, &schemaSyncHalvingConfigV0)
	if err != nil {
		panic("invalid embedded json for SyncHalvingConfigV0")
	}
	return schemaSyncHalvingConfigV0
}

func ParsedSearcherConfigV0() interface{} {
	if schemaSearcherConfigV0 != nil {
		return schemaSearcherConfigV0
	}
	err := json.Unmarshal(textSearcherConfigV0, &schemaSearcherConfigV0)
	if err != nil {
		panic("invalid embedded json for SearcherConfigV0")
	}
	return schemaSearcherConfigV0
}

func ParsedSecurityConfigV0() interface{} {
	if schemaSecurityConfigV0 != nil {
		return schemaSecurityConfigV0
	}
	err := json.Unmarshal(textSecurityConfigV0, &schemaSecurityConfigV0)
	if err != nil {
		panic("invalid embedded json for SecurityConfigV0")
	}
	return schemaSecurityConfigV0
}

func ParsedSharedFSConfigV0() interface{} {
	if schemaSharedFSConfigV0 != nil {
		return schemaSharedFSConfigV0
	}
	err := json.Unmarshal(textSharedFSConfigV0, &schemaSharedFSConfigV0)
	if err != nil {
		panic("invalid embedded json for SharedFSConfigV0")
	}
	return schemaSharedFSConfigV0
}

func ParsedTensorboardStorageConfigV0() interface{} {
	if schemaTensorboardStorageConfigV0 != nil {
		return schemaTensorboardStorageConfigV0
	}
	err := json.Unmarshal(textTensorboardStorageConfigV0, &schemaTensorboardStorageConfigV0)
	if err != nil {
		panic("invalid embedded json for TensorboardStorageConfigV0")
	}
	return schemaTensorboardStorageConfigV0
}

func ParsedTestRootV0() interface{} {
	if schemaTestRootV0 != nil {
		return schemaTestRootV0
	}
	err := json.Unmarshal(textTestRootV0, &schemaTestRootV0)
	if err != nil {
		panic("invalid embedded json for TestRootV0")
	}
	return schemaTestRootV0
}

func ParsedTestSubV0() interface{} {
	if schemaTestSubV0 != nil {
		return schemaTestSubV0
	}
	err := json.Unmarshal(textTestSubV0, &schemaTestSubV0)
	if err != nil {
		panic("invalid embedded json for TestSubV0")
	}
	return schemaTestSubV0
}

func ParsedTestUnionAV0() interface{} {
	if schemaTestUnionAV0 != nil {
		return schemaTestUnionAV0
	}
	err := json.Unmarshal(textTestUnionAV0, &schemaTestUnionAV0)
	if err != nil {
		panic("invalid embedded json for TestUnionAV0")
	}
	return schemaTestUnionAV0
}

func ParsedTestUnionBV0() interface{} {
	if schemaTestUnionBV0 != nil {
		return schemaTestUnionBV0
	}
	err := json.Unmarshal(textTestUnionBV0, &schemaTestUnionBV0)
	if err != nil {
		panic("invalid embedded json for TestUnionBV0")
	}
	return schemaTestUnionBV0
}

func ParsedTestUnionV0() interface{} {
	if schemaTestUnionV0 != nil {
		return schemaTestUnionV0
	}
	err := json.Unmarshal(textTestUnionV0, &schemaTestUnionV0)
	if err != nil {
		panic("invalid embedded json for TestUnionV0")
	}
	return schemaTestUnionV0
}

func schemaBytesMap() map[string][]byte {
	if cachedSchemaBytesMap != nil {
		return cachedSchemaBytesMap
	}
	var url string
	cachedSchemaBytesMap = map[string][]byte{}
	url = "http://determined.ai/schemas/expconf/v0/bind-mount.json"
	cachedSchemaBytesMap[url] = textBindMountV0
	url = "http://determined.ai/schemas/expconf/v0/check-data-layer-cache.json"
	cachedSchemaBytesMap[url] = textCheckDataLayerCacheV0
	url = "http://determined.ai/schemas/expconf/v0/check-epoch-not-used.json"
	cachedSchemaBytesMap[url] = textCheckEpochNotUsedV0
	url = "http://determined.ai/schemas/expconf/v0/check-global-batch-size.json"
	cachedSchemaBytesMap[url] = textCheckGlobalBatchSizeV0
	url = "http://determined.ai/schemas/expconf/v0/check-grid-hyperparameter.json"
	cachedSchemaBytesMap[url] = textCheckGridHyperparameterV0
	url = "http://determined.ai/schemas/expconf/v0/check-positive-length.json"
	cachedSchemaBytesMap[url] = textCheckPositiveLengthV0
	url = "http://determined.ai/schemas/expconf/v0/checkpoint-storage.json"
	cachedSchemaBytesMap[url] = textCheckpointStorageConfigV0
	url = "http://determined.ai/schemas/expconf/v0/data-layer-gcs.json"
	cachedSchemaBytesMap[url] = textGCSDataLayerConfigV0
	url = "http://determined.ai/schemas/expconf/v0/data-layer-s3.json"
	cachedSchemaBytesMap[url] = textS3DataLayerConfigV0
	url = "http://determined.ai/schemas/expconf/v0/data-layer-shared-fs.json"
	cachedSchemaBytesMap[url] = textSharedFSDataLayerConfigV0
	url = "http://determined.ai/schemas/expconf/v0/data-layer.json"
	cachedSchemaBytesMap[url] = textDataLayerConfigV0
	url = "http://determined.ai/schemas/expconf/v0/environment-image-map.json"
	cachedSchemaBytesMap[url] = textEnvironmentImageMapV0
	url = "http://determined.ai/schemas/expconf/v0/environment-image.json"
	cachedSchemaBytesMap[url] = textEnvironmentImageV0
	url = "http://determined.ai/schemas/expconf/v0/environment-variables-map.json"
	cachedSchemaBytesMap[url] = textEnvironmentVariablesMapV0
	url = "http://determined.ai/schemas/expconf/v0/environment-variables.json"
	cachedSchemaBytesMap[url] = textEnvironmentVariablesV0
	url = "http://determined.ai/schemas/expconf/v0/environment.json"
	cachedSchemaBytesMap[url] = textEnvironmentConfigV0
	url = "http://determined.ai/schemas/expconf/v0/experiment.json"
	cachedSchemaBytesMap[url] = textExperimentConfigV0
	url = "http://determined.ai/schemas/expconf/v0/gcs.json"
	cachedSchemaBytesMap[url] = textGCSConfigV0
	url = "http://determined.ai/schemas/expconf/v0/hdfs.json"
	cachedSchemaBytesMap[url] = textHDFSConfigV0
	url = "http://determined.ai/schemas/expconf/v0/hyperparameter-categorical.json"
	cachedSchemaBytesMap[url] = textCategoricalHyperparameterV0
	url = "http://determined.ai/schemas/expconf/v0/hyperparameter-const.json"
	cachedSchemaBytesMap[url] = textConstHyperparameterV0
	url = "http://determined.ai/schemas/expconf/v0/hyperparameter-double.json"
	cachedSchemaBytesMap[url] = textDoubleHyperparameterV0
	url = "http://determined.ai/schemas/expconf/v0/hyperparameter-int.json"
	cachedSchemaBytesMap[url] = textIntHyperparameterV0
	url = "http://determined.ai/schemas/expconf/v0/hyperparameter-log.json"
	cachedSchemaBytesMap[url] = textLogHyperparameterV0
	url = "http://determined.ai/schemas/expconf/v0/hyperparameter.json"
	cachedSchemaBytesMap[url] = textHyperparameterV0
	url = "http://determined.ai/schemas/expconf/v0/hyperparameters.json"
	cachedSchemaBytesMap[url] = textHyperparametersV0
	url = "http://determined.ai/schemas/expconf/v0/internal.json"
	cachedSchemaBytesMap[url] = textInternalConfigV0
	url = "http://determined.ai/schemas/expconf/v0/kerberos.json"
	cachedSchemaBytesMap[url] = textKerberosConfigV0
	url = "http://determined.ai/schemas/expconf/v0/length.json"
	cachedSchemaBytesMap[url] = textLengthV0
	url = "http://determined.ai/schemas/expconf/v0/optimizations.json"
	cachedSchemaBytesMap[url] = textOptimizationsConfigV0
	url = "http://determined.ai/schemas/expconf/v0/reproducibility.json"
	cachedSchemaBytesMap[url] = textReproducibilityConfigV0
	url = "http://determined.ai/schemas/expconf/v0/resources.json"
	cachedSchemaBytesMap[url] = textResourcesConfigV0
	url = "http://determined.ai/schemas/expconf/v0/s3.json"
	cachedSchemaBytesMap[url] = textS3ConfigV0
	url = "http://determined.ai/schemas/expconf/v0/searcher-adaptive-asha.json"
	cachedSchemaBytesMap[url] = textAdaptiveASHAConfigV0
	url = "http://determined.ai/schemas/expconf/v0/searcher-adaptive-simple.json"
	cachedSchemaBytesMap[url] = textAdaptiveSimpleConfigV0
	url = "http://determined.ai/schemas/expconf/v0/searcher-adaptive.json"
	cachedSchemaBytesMap[url] = textAdaptiveConfigV0
	url = "http://determined.ai/schemas/expconf/v0/searcher-async-halving.json"
	cachedSchemaBytesMap[url] = textAsyncHalvingConfigV0
	url = "http://determined.ai/schemas/expconf/v0/searcher-grid.json"
	cachedSchemaBytesMap[url] = textGridConfigV0
	url = "http://determined.ai/schemas/expconf/v0/searcher-pbt.json"
	cachedSchemaBytesMap[url] = textPBTConfigV0
	url = "http://determined.ai/schemas/expconf/v0/searcher-random.json"
	cachedSchemaBytesMap[url] = textRandomConfigV0
	url = "http://determined.ai/schemas/expconf/v0/searcher-single.json"
	cachedSchemaBytesMap[url] = textSingleConfigV0
	url = "http://determined.ai/schemas/expconf/v0/searcher-sync-halving.json"
	cachedSchemaBytesMap[url] = textSyncHalvingConfigV0
	url = "http://determined.ai/schemas/expconf/v0/searcher.json"
	cachedSchemaBytesMap[url] = textSearcherConfigV0
	url = "http://determined.ai/schemas/expconf/v0/security.json"
	cachedSchemaBytesMap[url] = textSecurityConfigV0
	url = "http://determined.ai/schemas/expconf/v0/shared-fs.json"
	cachedSchemaBytesMap[url] = textSharedFSConfigV0
	url = "http://determined.ai/schemas/expconf/v0/tensorboard-storage.json"
	cachedSchemaBytesMap[url] = textTensorboardStorageConfigV0
	url = "http://determined.ai/schemas/expconf/v0/test-root.json"
	cachedSchemaBytesMap[url] = textTestRootV0
	url = "http://determined.ai/schemas/expconf/v0/test-sub.json"
	cachedSchemaBytesMap[url] = textTestSubV0
	url = "http://determined.ai/schemas/expconf/v0/test-union-a.json"
	cachedSchemaBytesMap[url] = textTestUnionAV0
	url = "http://determined.ai/schemas/expconf/v0/test-union-b.json"
	cachedSchemaBytesMap[url] = textTestUnionBV0
	url = "http://determined.ai/schemas/expconf/v0/test-union.json"
	cachedSchemaBytesMap[url] = textTestUnionV0
	return cachedSchemaBytesMap
}
