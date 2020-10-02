"""
A one-variable proportional model that initializes its weight to zero, and whose data and labels
are always just 1.

Useful for testing metrics or gradient updates.
"""

import tensorflow as tf
import os
import io
import numpy as np

from determined import estimator
from determined import tensorboard
from PIL import Image


def tf_data_loader(batch_size, length):
    one_ds = tf.data.Dataset.from_tensor_slices(tf.constant([1], tf.float32))

    data = one_ds.repeat(length).batch(batch_size)
    label = one_ds.repeat(length).batch(batch_size)

    return tf.data.Dataset.zip(({"x": data}, label))


def make_model_fn(feature_columns, optimizer):

    def custom_model_fn(features, labels, mode):
        input_layer = tf.compat.v1.feature_column.input_layer(features, feature_columns)
        dense = tf.compat.v1.layers.Dense(
            units=1, use_bias=False, kernel_initializer=tf.zeros_initializer(), name="my_dense",
        )
        output_layer = dense(input_layer)
        predictions = tf.squeeze(output_layer, 1)

        if mode == tf.estimator.ModeKeys.PREDICT:
            return tf.estimator.EstimatorSpec(mode, predictions=predictions)

        loss = tf.losses.mean_squared_error(labels, predictions)

        if mode == tf.estimator.ModeKeys.EVAL:
            mae = tf.metrics.mean_absolute_error(labels, predictions)

            return tf.estimator.EstimatorSpec(
                mode,
                loss=loss,
                eval_metric_ops={"mae": mae},
            )

        if mode == tf.estimator.ModeKeys.TRAIN:
            train_op = optimizer.minimize(
                loss, global_step=tf.compat.v1.train.get_global_step()
            )
            return tf.estimator.EstimatorSpec(
               mode,
               loss=loss,
               train_op=train_op,
            )

    return custom_model_fn

# class MyEvalHook(tf.compat.v1.train.SessionRunHook):
class MyEvalHook(estimator.RunHook):
    """
    As a trivial example, MyEvalHook just writes a single scalar metric directly to the Determined
    TensorBoard directory.  The scalar value happens to be the weight of the one-variable model.
    """
    def begin(self):
        self.weight_in_my_dense_kernel = None
        self.global_step = None

    def before_run(self, run_contex):
        print("CONTEXT:")
        print(run_contex._original_args)
        if self.weight_in_my_dense_kernel is None:
            return tf.estimator.SessionRunArgs(
                fetches={"weight": "my_dense/kernel:0", "step": "global_step:0"}
            )

    def after_run(self, run_contex, run_values):
        if run_values.results is not None:
            self.weight_in_my_dense_kernel = run_values.results["weight"][0, 0]
            self.global_step = run_values.results["step"]

        w, h = 512, 512
        data = np.zeros((h, w, 3), dtype=np.uint8)
        data[0:256, 0:256] = [255, 0, 0] # red patch in upper left
        img = Image.fromarray(data, 'RGB')
        png_buffer = io.BytesIO()
        img.save(png_buffer, "png")
        png_encoded = png_buffer.getvalue()
        png_buffer.close()

        from tensorflow.core.util import event_pb2
        with tf.compat.v1.summary.FileWriter("/tmp/tensorboard/") as writer:
            summary = tf.Summary()
            summary_value = summary.value.add()
            summary_value.tag = "weight"
            summary_value.simple_value = self.weight_in_my_dense_kernel
            # writer.add_summary(summary)
            event = event_pb2.Event(summary=summary)
            event.step = self.global_step
            writer.add_event(event)

            summary_image = tf.Summary.Image(height=h, width=w, colorspace=4, encoded_image_string=png_encoded)
            summary2 = tf.Summary(value=[tf.Summary.Value(tag="my_image", image=summary_image)])
            writer.add_summary(summary2, self.global_step)

        w, h = 512, 512
        data = np.zeros((h, w, 3), dtype=np.uint8)
        data[0:256, 0:256] = [0, 255, 0] # red patch in upper left
        img = Image.fromarray(data, 'RGB')
        png_buffer = io.BytesIO()
        img.save(png_buffer, "png")
        png_encoded = png_buffer.getvalue()
        png_buffer.close()

        with tf.compat.v1.summary.FileWriter("/tmp/tensorboard/eval") as writer:
            summary_image = tf.Summary.Image(height=h, width=w, colorspace=4, encoded_image_string=png_encoded)
            summary2 = tf.Summary(value=[tf.Summary.Value(tag="my_image", image=summary_image)])
            writer.add_summary(summary2, self.global_step)

        # w, h = 512, 512
        # data = np.zeros((h, w, 3), dtype=np.uint8)
        # data[0:256, 0:256] = [0, 255, 0] # red patch in upper left
        # img = Image.fromarray(data, 'RGB')
        # png_buffer = io.BytesIO()
        # img.save(png_buffer, "png")
        # png_encoded = png_buffer.getvalue()
        # png_buffer.close()

        # with tf.compat.v1.summary.FileWriter("/tmp/tensorboard/eval") as writer:
        #     summary = tf.Summary()
        #     summary_value = summary.value.add()
        #     summary_value.tag = "weight"
        #     summary_value.simple_value = self.weight_in_my_dense_kernel
        #     # writer.add_summary(summary)
        #     event = event_pb2.Event(summary=summary)
        #     event.step = self.global_step
        #     writer.add_event(event)

        #     summary_image = tf.Summary.Image(height=h, width=w, colorspace=4, encoded_image_string=png_encoded)
        #     summary2 = tf.Summary(value=[tf.Summary.Value(tag="my_image", image=summary_image)])
        #     writer.add_summary(summary2, self.global_step)

    def end(self, session):
        # The key is to use this determined.tensorboard.get_base_path() call to find out where
        # you can write tensorboard metrics so that they will be persisted.  This is currently an
        # internal API, so the API may change, but the behavior is not going to go away.  We are
        # currently in API discussions as to how we want to make the public API for this behavior.
        # tensorflow_dir = tensorboard.get_base_path({})

        # There's probably a much easier way to write metrics to a directory but I'm honestly not
        # very good with tensorboard, so here goes:
        from tensorflow.core.util import event_pb2
        with tf.compat.v1.summary.FileWriter("/tmp/tensorboard") as writer:
            summary = tf.Summary()
            summary_value = summary.value.add()
            summary_value.tag = "weight"
            summary_value.simple_value = self.weight_in_my_dense_kernel
            # writer.add_summary(summary)
            event = event_pb2.Event(summary=summary)
            event.step = self.global_step
            writer.add_event(event)


class MyLinearEstimator(estimator.EstimatorTrial):
    def __init__(self, context: estimator.EstimatorTrialContext) -> None:
        self.context = context
        self.hparams = context.get_hparams()
        self.batch_size = self.context.get_per_slot_batch_size()

        self.dense = None

    def build_estimator(self) -> tf.estimator.Estimator:
        feature_columns=[tf.feature_column.numeric_column("x", shape=(), dtype=tf.int64)]
        optimizer = tf.compat.v1.train.GradientDescentOptimizer(
            learning_rate=self.hparams["learning_rate"],
        )
        optimizer = self.context.wrap_optimizer(optimizer)

        estimator = tf.compat.v1.estimator.Estimator(
            model_fn=make_model_fn(feature_columns, optimizer)
        )

        return estimator

    def build_train_spec(self) -> tf.estimator.TrainSpec:
        def fn():
            ds = tf_data_loader(self.context.get_per_slot_batch_size(), 100)
            return self.context.wrap_dataset(ds)

        return tf.estimator.TrainSpec(fn)

    def build_validation_spec(self) -> tf.estimator.EvalSpec:
        def fn():
            ds = tf_data_loader(self.context.get_per_slot_batch_size(), 10)
            return self.context.wrap_dataset(ds)

        return tf.estimator.EvalSpec(fn, hooks=[MyEvalHook()])
