/*
 * Copyright (c) 2013 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

#include <stdint.h>
#include <stdlib.h>

// GObject Type Casting
static GObject *
toGObject(void *p)
{
	return (G_OBJECT(p));
}

/* Wrapper to avoid variable arg list */
static void
_g_object_set_one(gpointer object, const gchar *property_name, void *val)
{
	g_object_set(object, property_name, *(gpointer **)val, NULL);
}

/* Wrapper to avoid variable arg list */
static void
_g_signal_emit_by_name_one(gpointer instance, const gchar *detailed_signal)
{
	g_signal_emit_by_name(instance, detailed_signal);
}

typedef struct {
	char            *name;
	int             func_n;
	void            *target;
	uintptr_t       *args;
	int             nargs;
	gboolean        ret;
	gulong          id;
} cbinfo;

/* Wrapper that runs the Go closure for a given context */
extern void _go_glib_callback(cbinfo *cbi);

/* Set up the args and call the Go closure wrapper */
static gboolean
_glib_callback(void *data, ...) {
	va_list         ap;
	cbinfo          *cbi = (cbinfo *)data;
	int             i;

	cbi->args = calloc(cbi->nargs, sizeof(void *));
	va_start(ap, data);
	for (i = 0; i < cbi->nargs; ++i)
		cbi->args[i] = va_arg(ap, uintptr_t);
	va_end(ap);

	_go_glib_callback(cbi);
	free(cbi->args);
	return (cbi->ret);
}

static void
cbinfo_free(gpointer data, GClosure *closure)
{
	free((cbinfo *)data);
}

static cbinfo *
_g_signal_connect(void *obj, gchar *name, int func_n)
{
	GSignalQuery    query;
	guint           sig_id;
	cbinfo          *cbi;

	sig_id = g_signal_lookup(name, G_OBJECT_TYPE(obj));
	g_signal_query(sig_id, &query);
	cbi = (cbinfo *)malloc(sizeof(cbinfo));
	cbi->name = g_strdup(name);
	cbi->func_n = func_n;
	cbi->args = NULL;
	cbi->target = obj;
	cbi->nargs = query.n_params;
	cbi->id = g_signal_connect_data((gpointer)obj, name,
	    G_CALLBACK(_glib_callback), cbi, cbinfo_free, G_CONNECT_SWAPPED);
	return (cbi);
}

static uintptr_t
cbinfo_get_arg(cbinfo *cbi, int n) {
	return (cbi->args[n]);
}

static gulong
cbinfo_get_id(cbinfo *cbi)
{
	return (cbi->id);
}

typedef struct {
	int		func_n;
	gboolean	ret;
	guint		id;
} idleinfo;

/* Call to Go to nil this context to free memory next garbage collection */
extern void _go_nil_unused_idle_ctx(int n);

/*
 * Called after a function is removed from the main loop context to free the
 * idle context.
 */
static void
idleinfo_free(gpointer data)
{
	idleinfo	*idl = (idleinfo *)data;

	_go_nil_unused_idle_ctx(idl->func_n);
	free(idl);
}

/* Call to Go to run func in the main loop context */
extern void _go_glib_idle_fn(idleinfo *idl);

static gboolean
_g_idle_run(gpointer user_data)
{
	idleinfo	*idl = (idleinfo *)user_data;

	_go_glib_idle_fn(idl);
	return (idl->ret);
}

/*
 * Create idleinfo context and add _g_idle_run and its context to run
 * in the GTK main loop during idle state
 */
static idleinfo *
_g_idle_add(int func_n)
{
	idleinfo	*idl;

	idl = (idleinfo *)malloc(sizeof(idleinfo));
	idl->func_n = func_n;
	idl->id = g_idle_add_full(G_PRIORITY_DEFAULT_IDLE, _g_idle_run,
	    (gpointer)idl, idleinfo_free);
	return (idl);
}

/*
 * GValue
 */

static GValue *
_g_value_alloc()
{
	GValue		value = G_VALUE_INIT;

	/* Can't do a sizeof(GValue) */
	return ((GValue *)g_malloc0(sizeof(value)));
}

static GValue *
_g_value_init(GType g_type)
{
	GValue          *value;

	value = g_new0(GValue, 1);
	return (g_value_init(value, g_type));
}
