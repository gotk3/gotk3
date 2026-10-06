// Same copyright and license as the rest of the files in this project

/*
 * GAsyncReadyCallback
 */

extern void goAsyncReadyCallbacks(GObject *source_object, GAsyncResult *res,
                                  guintptr user_data);

static inline void _gotk3_goAsyncReadyCallbacks(GObject *source_object,
                                                GAsyncResult *res,
                                                gpointer user_data) {
  goAsyncReadyCallbacks(source_object, res, (guintptr)user_data);
}

static inline void _g_permission_acquire_async(GPermission *permission,
                                               GCancellable *cancellable,
                                               guintptr user_data) {
  g_permission_acquire_async(permission, cancellable,
                             _gotk3_goAsyncReadyCallbacks, (gpointer)user_data);
}

static inline void _g_permission_release_async(GPermission *permission,
                                               GCancellable *cancellable,
                                               guintptr user_data) {
  g_permission_release_async(permission, cancellable,
                             _gotk3_goAsyncReadyCallbacks, (gpointer)user_data);
}
