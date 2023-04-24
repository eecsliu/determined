import pathlib
import subprocess

import pytest
import tensorflow as tf
from packaging import version

official_examples = [
    (
        "tutorials/mnist_pytorch",
        "tutorials/mnist_pytorch/const.yaml",
    ),
    (
        "tutorials/fashion_mnist_tf_keras",
        "tutorials/fashion_mnist_tf_keras/const.yaml",
    ),
    (
        "tutorials/imagenet_pytorch",
        "tutorials/imagenet_pytorch/const_cifar.yaml",
    ),
    (
        "computer_vision/cifar10_pytorch",
        "computer_vision/cifar10_pytorch/const.yaml",
    ),
    (
        "computer_vision/mnist_estimator",
        "computer_vision/mnist_estimator/const.yaml",
    ),
    (
        "computer_vision/cifar10_tf_keras",
        "computer_vision/cifar10_tf_keras/const.yaml",
    ),
    (
        "computer_vision/iris_tf_keras",
        "computer_vision/iris_tf_keras/const.yaml",
    ),
    (
        "gan/gan_mnist_pytorch",
        "gan/gan_mnist_pytorch/const.yaml",
    ),
    (
        "gan/dcgan_tf_keras",
        "gan/dcgan_tf_keras/const.yaml",
    ),
    (
        "gan/pix2pix_tf_keras",
        "gan/pix2pix_tf_keras/const.yaml",
    ),
    (
        "decision_trees/gbt_titanic_estimator",
        "decision_trees/gbt_titanic_estimator/const.yaml",
    ),
    (
        "features/custom_reducers_mnist_pytorch",
        "features/custom_reducers_mnist_pytorch/const.yaml",
    ),
]


@pytest.mark.parametrize("model_def,config_file", official_examples)
def test_official(model_def: str, config_file: str) -> None:
    if (
        version.parse(tf.__version__) >= version.parse("2.11.0")
        and "gbt_titanic_estimator" in model_def
    ):
        pytest.skip(f"# TODO [MLG-442], see comment in {model_def}")
    examples_dir = pathlib.Path(__file__).parent.parent
    model_def_absolute = examples_dir.joinpath(model_def)
    config_file_absolute = examples_dir.joinpath(config_file)

    startup_hook = model_def_absolute.joinpath("startup-hook.sh")
    if startup_hook.exists():
        subprocess.check_output(("sh", str(startup_hook)))

    subprocess.check_output(
        (
            "det",
            "experiment",
            "create",
            "--local",
            "--test",
            str(config_file_absolute),
            str(model_def_absolute),
        )
    )
