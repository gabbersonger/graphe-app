// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

/**
 * A Logger records structured information about each call to its
 * Log, Debug, Info, Warn, and Error methods.
 * For each call, it creates a [Record] and passes it to a [Handler].
 * 
 * To create a new Logger, call [New] or a Logger method
 * that begins "With".
 * @module
 */

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "@wailsio/runtime";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as $models from "./models.js";

/**
 * Debug logs at [LevelDebug].
 */
export function Debug(msg: string, ...args: any[]): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1366249994, msg, args) as any;
    return $resultPromise;
}

/**
 * DebugContext logs at [LevelDebug] with the given context.
 */
export function DebugContext(msg: string, ...args: any[]): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1149598869, msg, args) as any;
    return $resultPromise;
}

/**
 * Enabled reports whether l emits log records at the given context and level.
 */
export function Enabled(level: $models.Level): Promise<boolean> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2968345048, level) as any;
    return $resultPromise;
}

/**
 * Error logs at [LevelError].
 */
export function Error(msg: string, ...args: any[]): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(744032651, msg, args) as any;
    return $resultPromise;
}

/**
 * ErrorContext logs at [LevelError] with the given context.
 */
export function ErrorContext(msg: string, ...args: any[]): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(807530258, msg, args) as any;
    return $resultPromise;
}

/**
 * Handler returns l's Handler.
 */
export function Handler(): Promise<$models.Handler> & { cancel(): void } {
    let $resultPromise = $Call.ByID(89678517) as any;
    return $resultPromise;
}

/**
 * Info logs at [LevelInfo].
 */
export function Info(msg: string, ...args: any[]): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(97676471, msg, args) as any;
    return $resultPromise;
}

/**
 * InfoContext logs at [LevelInfo] with the given context.
 */
export function InfoContext(msg: string, ...args: any[]): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3150713246, msg, args) as any;
    return $resultPromise;
}

/**
 * Log emits a log record with the current time and the given level and message.
 * The Record's Attrs consist of the Logger's attributes followed by
 * the Attrs specified by args.
 * 
 * The attribute arguments are processed as follows:
 *   - If an argument is an Attr, it is used as is.
 *   - If an argument is a string and this is not the last argument,
 *     the following argument is treated as the value and the two are combined
 *     into an Attr.
 *   - Otherwise, the argument is treated as a value with key "!BADKEY".
 */
export function Log(level: $models.Level, msg: string, ...args: any[]): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2951997251, level, msg, args) as any;
    return $resultPromise;
}

/**
 * LogAttrs is a more efficient version of [Logger.Log] that accepts only Attrs.
 */
export function LogAttrs(level: $models.Level, msg: string, ...attrs: $models.Attr[]): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2232660861, level, msg, attrs) as any;
    return $resultPromise;
}

/**
 * Warn logs at [LevelWarn].
 */
export function Warn(msg: string, ...args: any[]): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1007172339, msg, args) as any;
    return $resultPromise;
}

/**
 * WarnContext logs at [LevelWarn] with the given context.
 */
export function WarnContext(msg: string, ...args: any[]): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1614796026, msg, args) as any;
    return $resultPromise;
}

/**
 * With returns a Logger that includes the given attributes
 * in each output operation. Arguments are converted to
 * attributes as if by [Logger.Log].
 */
export function With(...args: any[]): Promise<$models.Logger | null> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3480014187, args) as any;
    let $typingPromise = $resultPromise.then(($result) => {
        return $$createType1($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

/**
 * WithGroup returns a Logger that starts a group, if name is non-empty.
 * The keys of all attributes added to the Logger will be qualified by the given
 * name. (How that qualification happens depends on the [Handler.WithGroup]
 * method of the Logger's Handler.)
 * 
 * If name is empty, WithGroup returns the receiver.
 */
export function WithGroup(name: string): Promise<$models.Logger | null> & { cancel(): void } {
    let $resultPromise = $Call.ByID(121027970, name) as any;
    let $typingPromise = $resultPromise.then(($result) => {
        return $$createType1($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

// Private type creation functions
const $$createType0 = $models.Logger.createFrom;
const $$createType1 = $Create.Nullable($$createType0);