
export function secondsToDateString(seconds: number): string {
    const t = new Date(1970, 0, 1); // Epoch
    t.setSeconds(seconds);
    return t.toISOString().split('T')[0];
}

export function secondsToDateTimeString(seconds: number): string {
    const t = new Date(1970, 0, 1); // Epoch
    t.setSeconds(seconds);
    return t.toISOString();
}
