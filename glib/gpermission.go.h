// Same copyright and license as the rest of the files in this project

/*
 * GAsyncReadyCallback
 */

extern void goAsyncReadyCallbacks(GObject *source_object, GAsyncResult *res,
                                  gpointer user_data);

static inline void _g_permission_acquire_async(GPermission *permission,
                                               GCancellable *cancellable,
                                               gpointer user_data) {
  g_permission_acquire_async(permission, cancellable,
                             (GAsyncReadyCallback)(goAsyncReadyCallbacks),
                             user_data);
}

static inline void _g_permission_release_async(GPermission *permission,
                                               GCancellable *cancellable,
                                               gpointer user_data) {
  g_permission_release_async(permission, cancellable,
                             (GAsyncReadyCallback)(goAsyncReadyCallbacks),
                             user_data);
}
